package main

import (
	"context"
	"net"

	"github.com/GP-Hacks/notifications/internal/service_provider"
	"github.com/GP-Hacks/notifications/internal/utils/logger"
	proto "github.com/GP-Hacks/proto/pkg/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger.Initialize(logger.Config{
		LogLevel:    "debug",
		Development: true,
		OutputPaths: []string{"stdout", "app.log"},
	})
	defer logger.Sync()
	logger.Info(
		"Application started",
	)

	serviceProvider := service_provider.NewServiceProvider()

	defer func() {
		if err := serviceProvider.MongoClient().Disconnect(context.Background()); err != nil {
			logger.Error(
				"Failed disconnect to MongoDB",
				zap.Error(err),
			)
		}
		if err := serviceProvider.RabbitmqConnection().Close(); err != nil {
			logger.Error(
				"Failed close RabbitMQ connection",
				zap.Error(err),
			)
		}
	}()

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(grpcServer)

	proto.RegisterNotificationsServer(grpcServer, serviceProvider.TokensController())

	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Fatal(
			"Failed start listen",
			zap.Error(err),
		)
	}

	err = grpcServer.Serve(list)
	if err != nil {
		logger.Fatal(
			"Failed serve grpc",
			zap.Error(err),
		)
	}
}
