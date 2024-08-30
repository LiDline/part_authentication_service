package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port          string `env:"PORT" env-default:"5000"`
	DbUrl         string `env:"DATABASE_URL" env-required:"true"`
	Secret        string `env:"SECRET_KEY" env-required:"true"`
	Email         string `env:"EMAIL"`
	PasswordEmail string `env:"PASSWORD_EMAIL"`
}

var Port string
var DbUrl string
var Secret string
var Email string
var PasswordEmail string

func MustLoadEnv() Config {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)

	if err != nil {
		log.Fatalf("Some environment is not set: %v", err)
	}

	Port = cfg.Port
	DbUrl = cfg.DbUrl
	Secret = cfg.Secret
	Email = cfg.Email
	PasswordEmail = cfg.PasswordEmail

	return cfg
}
