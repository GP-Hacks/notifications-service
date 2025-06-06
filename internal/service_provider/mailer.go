package service_provider

import "github.com/GP-Hacks/notifications/internal/infrastructure/mailer"

func (s *ServiceProvider) Mailer() *mailer.Mailer {
	if s.mailer == nil {
		s.mailer = mailer.NewMailer()
	}

	return s.mailer
}
