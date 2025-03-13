package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	GRPC struct {
		Port        string `env:"GRPC_PORT"`
		GatewayPort string `env:"GRPC_GATEWAY_PORT"`
	}
	PG struct {
		URL string `env:"POSTGRES_URL"`
	}
}

func New() (*Config, error) {
	cfg := &Config{}

	cfg.GRPC.Port = os.Getenv("GRPC_PORT")
	cfg.GRPC.GatewayPort = os.Getenv("GRPC_GATEWAY_PORT")

	cfg.PG.URL = os.Getenv("POSTGRES_URL")

	return cfg, nil
}
