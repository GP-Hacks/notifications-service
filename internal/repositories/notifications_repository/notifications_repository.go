package notifications_repository

import (
	"firebase.google.com/go/messaging"
	"github.com/GP-Hacks/notifications/internal/utils/logger"
	"go.uber.org/zap"
)

type NotificationsRepository struct {
	logger *zap.Logger
	client *messaging.Client
}

func NewNotificationsRepository(client *messaging.Client) *NotificationsRepository {
	lg := logger.With(zap.String("from", "NotificationsRepository"))
	return &NotificationsRepository{
		client: client,
		logger: lg,
	}
}
