package service_provider

import (
	"context"

	"github.com/GP-Hacks/notifications/internal/config"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *ServiceProvider) MongoClient() *mongo.Client {
	if s.mongoClient == nil {
		clientOpt := options.Client().ApplyURI(config.Cfg.MongoDB.Path)
		cl, err := mongo.Connect(context.Background(), clientOpt)
		if err != nil {
			log.Fatal().Msg("Failed connect to MongoDB")
		}

		s.mongoClient = cl
	}

	return s.mongoClient
}

func (s *ServiceProvider) MongoCollection() *mongo.Collection {
	if s.mongoCollection == nil {
		s.mongoCollection = s.MongoClient().Database(config.Cfg.MongoDB.Name).Collection(config.Cfg.MongoDB.Collection)
	}

	return s.mongoCollection
}
