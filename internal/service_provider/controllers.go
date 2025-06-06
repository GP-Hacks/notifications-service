package service_provider

import (
	"github.com/GP-Hacks/notifications/internal/controllers/grpc"
	"github.com/GP-Hacks/notifications/internal/controllers/rabbitmq"
)

func (s *ServiceProvider) NotificationsController() *rabbitmq.NotificationsController {
	if s.notificationsController == nil {
		s.notificationsController = rabbitmq.NewNotificationsController(s.RabbitmqConnection(), s.NotificationsService())
	}

	return s.notificationsController
}

func (s *ServiceProvider) TokensController() *grpc.TokensController {
	if s.tokensController == nil {
		s.tokensController = grpc.NewTokensController(s.NotificationsService())
	}

	return s.tokensController
}

func (s *ServiceProvider) EmailController() *rabbitmq.EmailController {
	if s.emailController == nil {
		s.emailController = rabbitmq.NewEmailController(s.RabbitmqConnection(), s.EmailService())
	}

	return s.emailController
}
