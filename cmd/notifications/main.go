package main

import (
	"context"
	"net"

	"github.com/GP-Hacks/notifications/internal/config"
	"github.com/GP-Hacks/notifications/internal/service_provider"
	"github.com/GP-Hacks/notifications/internal/utils/logger"
	proto "github.com/GP-Hacks/proto/pkg/api/notifications"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.LoadConfig("./config")
	logger.SetupLogger()

	serviceProvider := service_provider.NewServiceProvider()

	log.Info().Msg("Applications started")
	if config.Cfg.Logging.IsProduction {
		log.Info().Msg("Started in production mode")
	} else {
		log.Info().Msg("Started in debug mode")
	}

	defer func() {
		if err := serviceProvider.MongoClient().Disconnect(context.Background()); err != nil {
			log.Error().Msg("Failed disconnect to MongoDB")
		}
		if err := serviceProvider.RabbitmqConnection().Close(); err != nil {
			log.Error().Msg("Failed close RabbitMQ connection")
		}
	}()

	if err := serviceProvider.NotificationsController().StartConsumer(); err != nil {
		log.Error().Msg(err.Error())
	}

	if err := serviceProvider.EmailController().StartConsumer(); err != nil {
		log.Error().Msg(err.Error())
	}

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(grpcServer)

	proto.RegisterNotificationsServer(grpcServer, serviceProvider.TokensController())

	list, err := net.Listen("tcp", ":"+config.Cfg.Grpc.Port)
	if err != nil {
		log.Fatal().Msg("Failed start listen port")
	}

	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatal().Msg("Failed serve grpc")
	}
}
