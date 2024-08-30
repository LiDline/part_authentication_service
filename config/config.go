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

var Port string
var DbUrl string
var Secret string

func MustLoadEnv() Config {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)

	if err != nil {
		log.Fatalf("Some environment is not set: %v", err)
	}

	Port = cfg.Port
	DbUrl = cfg.DbUrl
	Secret = cfg.Secret

	return cfg
}
