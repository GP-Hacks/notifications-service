package service_provider

import (
	"github.com/GP-Hacks/notifications/internal/services/email_service"
	notification_service "github.com/GP-Hacks/notifications/internal/services/notifications_service"
)

func (s *ServiceProvider) NotificationsService() *notification_service.NotificationsService {
	if s.notificationsService == nil {
		s.notificationsService = notification_service.NewNotificationsService(s.TokensRepository(), s.NotificationsRepository())
	}

	return s.notificationsService
}

func (s *ServiceProvider) EmailService() *email_service.EmailService {
	if s.emailService == nil {
		s.emailService = email_service.NewEmailService(s.Mailer())
	}

	return s.emailService
}
