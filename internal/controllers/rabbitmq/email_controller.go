package rabbitmq

import (
	"encoding/json"

	"github.com/GP-Hacks/notifications/internal/config"
	"github.com/GP-Hacks/notifications/internal/models"
	"github.com/GP-Hacks/notifications/internal/services/email_service"
	"github.com/streadway/amqp"
)

type EmailController struct {
	connection   *amqp.Connection
	emailService *email_service.EmailService
}

func NewEmailController(conn *amqp.Connection, service *email_service.EmailService) *EmailController {
	return &EmailController{
		connection:   conn,
		emailService: service,
	}
}

func (c *EmailController) StartConsumer() error {
	ch, err := c.connection.Channel()
	if err != nil {
		return err
	}

	if _, err := ch.QueueDeclare(
		config.Cfg.RabbitMQ.EmailQueue,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	msgs, err := ch.Consume(
		config.Cfg.RabbitMQ.EmailQueue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for {
			msg := <-msgs

			var mail models.Mail
			if err := json.Unmarshal(msg.Body, &mail); err != nil {
				continue
			}

			if mail.Header == "" || mail.Body == "" || mail.To == "" {
				continue
			}

			c.emailService.Send(&mail)
		}
	}()

	return nil
}
