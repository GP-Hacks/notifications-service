package service_provider

import (
	"context"

	"github.com/GP-Hacks/notifications/config"
	"github.com/GP-Hacks/notifications/internal/utils/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func (s *ServiceProvider) MongoClient() *mongo.Client {
	if s.mongoClient == nil {
		clientOpt := options.Client().ApplyURI(config.Cfg.MongoDBPath)
		cl, err := mongo.Connect(context.Background(), clientOpt)
		if err != nil {
			logger.Fatal("Failed connect to MongoDB", zap.Error(err))
		}

		s.mongoClient = cl
	}

	return s.mongoClient
}

func (s *ServiceProvider) MongoCollection() *mongo.Collection {
	if s.mongoCollection == nil {
		s.mongoCollection = s.MongoClient().Database(config.Cfg.MongoDBPath).Collection(config.Cfg.MongoDBCollection)
	}

	return s.mongoCollection
}
