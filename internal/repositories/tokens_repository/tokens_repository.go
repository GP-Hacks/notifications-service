package tokens_repository

import (
	"github.com/GP-Hacks/notifications/internal/utils/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type TokensRepository struct {
	collection *mongo.Collection
	logger     *zap.Logger
}

func NewTokensRepository(collection *mongo.Collection) *TokensRepository {
	lg := logger.With(zap.String("from", "TokensRepository"))
	return &TokensRepository{
		collection: collection,
		logger:     lg,
	}
}
