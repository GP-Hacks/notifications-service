package email_service

import "github.com/GP-Hacks/notifications/internal/models"

type (
	IMailer interface {
		Send(mail *models.Mail) error
	}

	EmailService struct {
		mailer IMailer
	}
)

func NewEmailService(m IMailer) *EmailService {
	return &EmailService{
		mailer: m,
	}
}
