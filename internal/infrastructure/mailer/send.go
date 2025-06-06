package mailer

import (
	"github.com/GP-Hacks/notifications/internal/config"
	"github.com/GP-Hacks/notifications/internal/models"
	"github.com/GP-Hacks/notifications/internal/services"
	"github.com/rs/zerolog/log"
	"gopkg.in/gomail.v2"
)

func (m *Mailer) Send(mail *models.Mail) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", config.Cfg.Mail.From)
	msg.SetHeader("To", mail.To)
	msg.SetHeader("Subject", mail.Header)
	msg.SetBody("text/html", mail.Body)

	d := gomail.NewDialer(config.Cfg.Mail.Host, config.Cfg.Mail.Port, config.Cfg.Mail.Username, config.Cfg.Mail.Password)
	if err := d.DialAndSend(msg); err != nil {
		log.Error().Msg(err.Error())

		return services.InternalServerError
	}

	return nil
}
