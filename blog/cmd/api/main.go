package main

import (
	"easyapi/blog/internal/app"
	"easyapi/blog/internal/config"
	"log"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatalf("can not initialize logger: %s", err)
	}

	cfg, err := config.New()

	if err != nil {
		log.Fatalf("can not initialize config: %s", err)
	}

	app.Run(logger, cfg)
}
