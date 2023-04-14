package config

import (
	"go.uber.org/zap"
	"hospital/src/internal/modules/app"
)

type Config struct {
	Secret string `envconfig:"SECRET"`

	DBUser string `envconfig:"DB_USER"`
	DBPass string `envconfig:"DB_PASS"`
	DBHost string `envconfig:"DB_HOST" default:"db"`
	DBPort string `envconfig:"DB_PORT" default:"5432"`
	DBName string `envconfig:"DB_NAME" default:"main"`

	AutoMigrate bool   `envconfig:"AUTO_MIGRATE" default:"false"`
	LogLevel    string `envconfig:"LOG_LEVEL" default:"info" validate:"oneof=debug info warn error dpanic panic fatal"`
}

func NewConfig(app app.App, logger *zap.Logger, logLevel zap.AtomicLevel) (Config, error) {
	config := Config{
		AutoMigrate: true,
		Secret:      "1",
	}

	logger.Info("Configurated", zap.Any("config", config))

	err := logLevel.UnmarshalText([]byte(config.LogLevel))
	if err != nil {
		return Config{}, err
	}

	return config, err
}
