package email_service

import "github.com/GP-Hacks/notifications/internal/models"

func (s *EmailService) Send(m *models.Mail) error {
	return s.mailer.Send(m)
}
