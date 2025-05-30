package service_provider

import (
	"context"
	"encoding/json"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/GP-Hacks/notifications/internal/config"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

func (s *ServiceProvider) FirebaseApp() *firebase.App {
	if s.firebaseApp == nil {
		creds := map[string]string{
			"type":                        "service_account",
			"project_id":                  config.Cfg.Firebase.ProjectId,
			"private_key_id":              config.Cfg.Firebase.PrivateKeyId,
			"private_key":                 config.Cfg.Firebase.PrivateKey,
			"client_email":                config.Cfg.Firebase.ClientEmail,
			"client_id":                   config.Cfg.Firebase.ClientId,
			"auth_uri":                    "https://accounts.google.com/o/oauth2/auth",
			"token_uri":                   "https://oauth2.googleapis.com/token",
			"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
			"client_x509_cert_url":        config.Cfg.Firebase.ClientX509CertUrl,
			"universe_domain":             "googleapis.com",
		}

		credentialsJSON, err := json.Marshal(creds)
		if err != nil {
			log.Fatal().Msg("Failed unmarshal credentials")
		}

		opt := option.WithCredentialsJSON(credentialsJSON)
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatal().Msg("Failed init Firebase App")
		}

		s.firebaseApp = app
	}

	return s.firebaseApp
}

func (s *ServiceProvider) MessagingClient() *messaging.Client {
	if s.messagingClient == nil {
		client, err := s.FirebaseApp().Messaging(context.Background())
		if err != nil {
			log.Fatal().Msg("Failed get MessagingClient from Firebase App")
		}

		s.messagingClient = client
	}

	return s.messagingClient
}
