package service_provider

import (
	"github.com/GP-Hacks/notifications/config"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

func (s *ServiceProvider) RabbitmqConnection() *amqp.Connection {
	if s.rabbitmqConnection == nil {
		conn, err := amqp.Dial(config.Cfg.RabbitMQAddress)
		if err != nil {
			log.Fatal().Msg("Failed connect to RabbitMQ")
		}

		s.rabbitmqConnection = conn
	}

	return s.rabbitmqConnection
}
