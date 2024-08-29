package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port   string `env:"PORT" env-default:"5000"`
	DbUrl  string `env:"DATABASE_URL" env-required:"true"`
	Secret string `env:"SECRET_KEY" env-required:"true"`
}

func MustLoadEnv() Config {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)

	if err != nil {
		log.Fatalf("DATABASE_URL environment variable is not set: %v", err)
	}

	return cfg
}
