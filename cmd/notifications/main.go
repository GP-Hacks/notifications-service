package main

import (
	"context"
	"net"

	"github.com/GP-Hacks/notifications/internal/service_provider"
	"github.com/GP-Hacks/notifications/internal/utils/logger"
	proto "github.com/GP-Hacks/proto/pkg/api/notifications"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	serviceProvider := service_provider.NewServiceProvider()
	logger.SetupLogger("http://vector:9880")

	log.Info().Msg("Applications started")

	defer func() {
		if err := serviceProvider.MongoClient().Disconnect(context.Background()); err != nil {
			log.Error().Msg("Failed disconnect to MongoDB")
		}
		if err := serviceProvider.RabbitmqConnection().Close(); err != nil {
			log.Error().Msg("Failed close RabbitMQ connection")
		}
	}()

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(grpcServer)

	proto.RegisterNotificationsServer(grpcServer, serviceProvider.TokensController())

	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal().Msg("Failed start listen port")
	}

	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatal().Msg("Failed serve grpc")
	}
}
