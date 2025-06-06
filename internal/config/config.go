package config

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Grpc struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"grpc"`

	RabbitMQ struct {
		Address            string `mapstructure:"address"`
		NotificationsQueue string `mapstructure:"notifications_queue"`
		EmailQueue         string `mapstructure:"email_queue"`
	} `mapstructure:"rabbitmq"`

	MongoDB struct {
		Name       string `mapstructure:"name"`
		Collection string `mapstructure:"collection"`
		Path       string `mapstructure:"path"`
	} `mapstructure:"mongodb"`

	Firebase struct {
		ProjectId         string `mapstructure:"projectId"`
		PrivateKeyId      string `mapstructure:"privateKeyId"`
		PrivateKey        string `mapstructure:"privateKey"`
		ClientEmail       string `mapstructure:"clientEmail"`
		ClientId          string `mapstructure:"clientId"`
		ClientX509CertUrl string `mapstructure:"clientX509CertUrl"`
	} `mapstructure:"firebase"`

	Logging struct {
		IsProduction bool   `mapstructure:"isProduction"`
		VectorURL    string `mapstructure:"vectorURL"`
	} `mapstructure:"logging"`

	Mail struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		From     string `mapstructure:"from"`
	} `mapstructure:"mail"`
}

var Cfg Config

func LoadConfig(path string) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)

	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	setDefaults(v)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Msg("Config not found, used env and default values")
		} else {
			log.Error().Msg("Failed read config file")
		}
	}

	for _, k := range v.AllKeys() {
		value := v.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			envVar := strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")
			envValue := os.Getenv(envVar)
			if envValue != "" {
				v.Set(k, envValue)
			}
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Error().Msg("Failed unmarshal config")
	}

	if err := validateConfig(&cfg); err != nil {
		log.Fatal().Msg("Not valid config")
	}

	Cfg = cfg
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("grpc.port", "50051")

	v.SetDefault("rabbitmq.address", "amqp://guest:guest@localhost:5672/")
	v.SetDefault("rabbitmq.notifications_queue", "tasks")
	v.SetDefault("rabbitmq.email_queue", "tasks")

	v.SetDefault("mongodb.name", "mydb")
	v.SetDefault("mongodb.collection", "documents")
	v.SetDefault("mongodb.path", "mongodb://localhost:27017")

	v.SetDefault("logging.isProduction", false)
	v.SetDefault("logging.vectorURL", "http://vector:9880")

	v.SetDefault("mail.host", "tasks")
	v.SetDefault("mail.port", "tasks")
	v.SetDefault("mail.username", "tasks")
	v.SetDefault("mail.password", "tasks")
	v.SetDefault("mail.from", "tasks")
}

func validateConfig(cfg *Config) error {
	if cfg.Firebase.ProjectId == "" {
		// return fmt.Errorf("firebase.projectId не задан")
	}

	return nil
}
