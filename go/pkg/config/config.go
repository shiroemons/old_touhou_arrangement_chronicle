package config

import (
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Env          string `env:"ENV" envDefault:"dev"`
	Port         string `env:"PORT" envDefault:"8080"`
	ConnectURL   string `env:"CONNECT_URL"`
	AllowOrigins string `env:"ALLOW_ORIGINS"`
}

func loadEnv() {
	e := os.Getenv("ENV")
	if e == "" || e == "dev" {
		err := godotenv.Load(".env.development.local")
		if err != nil {
			log.Fatal("Error loading .env.development.local file")
		}
	}
}

func Provider() Config {
	loadEnv()

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalln(err)
	}
	return *cfg
}
