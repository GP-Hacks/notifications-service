package main

import (
	"context"
	"log"
	"net"

	"github.com/GP-Hacks/notifications/internal/service_provider"
	proto "github.com/GP-Hacks/proto/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	serviceProvider := service_provider.NewServiceProvider()

	defer func() {
		if err := serviceProvider.MongoClient().Disconnect(context.Background()); err != nil {
			log.Printf("Failed to disconnect MongoDB: %v", err)
		}
		if err := serviceProvider.RabbitmqConnection().Close(); err != nil {
			log.Printf("Failed to close RabbitMQ connection: %v", err)
		}
	}()

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(grpcServer)

	proto.RegisterNotificationsServer(grpcServer, serviceProvider.TokensController())

	list, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	err = grpcServer.Serve(list)
	if err != nil {
		panic(err)
	}
}
