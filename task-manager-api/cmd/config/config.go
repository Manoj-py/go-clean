package config

import (
	"github.com/Manoj-py/backend-architecture/common/environment"
	"log"
	"log/slog"
)

type Config struct {
	PORT         string `env:"PORT"`
	POSTGRES_URL string `env:"POSTGRES_URL"`
}

func MustLoadConfig(env environment.LoadConfigI) *Config {
	config := Config{}
	if err := env.LoadEnv(&config); err != nil {
		slog.Error("couldn't load config err: ", err)
		log.Fatal(err)
		return nil
	}

	slog.Info("config loaded sucessfully..")
	return &config
}
