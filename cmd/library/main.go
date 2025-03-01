package main

import (
	"library/config"
	"library/internal/app"
	"log"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	var logger *zap.Logger

	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(logger, cfg)
}
