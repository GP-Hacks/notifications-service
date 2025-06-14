package service_provider

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/GP-Hacks/notifications/internal/controllers/grpc"
	"github.com/GP-Hacks/notifications/internal/controllers/rabbitmq"
	"github.com/GP-Hacks/notifications/internal/infrastructure/mailer"
	"github.com/GP-Hacks/notifications/internal/repositories/notifications_repository"
	"github.com/GP-Hacks/notifications/internal/repositories/tokens_repository"
	"github.com/GP-Hacks/notifications/internal/services/email_service"
	notification_service "github.com/GP-Hacks/notifications/internal/services/notifications_service"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

// ServiceProvider struct  
// Struct for provide service objects
type ServiceProvider struct {
	mailer                  *mailer.Mailer
	emailService            *email_service.EmailService
	emailController         *rabbitmq.EmailController
	notificationsController *rabbitmq.NotificationsController
	tokensController        *grpc.TokensController
	notificationsService    *notification_service.NotificationsService
	notificationsRepository *notifications_repository.NotificationsRepository
	tokensRepository        *tokens_repository.TokensRepository
	mongoCollection         *mongo.Collection
	mongoClient             *mongo.Client
	firebaseApp             *firebase.App
	messagingClient         *messaging.Client
	rabbitmqConnection      *amqp.Connection
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
