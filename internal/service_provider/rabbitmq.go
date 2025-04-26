package service_provider

import (
	"github.com/GP-Hacks/notifications/config"
	"github.com/GP-Hacks/notifications/internal/utils/logger"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func (s *ServiceProvider) RabbitmqConnection() *amqp.Connection {
	if s.rabbitmqConnection == nil {
		conn, err := amqp.Dial(config.Cfg.RabbitMQAddress)
		if err != nil {
			logger.Fatal("Failed connect to RabbitMQ", zap.Error(err))
		}

		s.rabbitmqConnection = conn
	}

	return s.rabbitmqConnection
}
