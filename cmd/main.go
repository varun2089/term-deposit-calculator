package main

import (
	"os"
	"term-deposit-calculator/internal/config"
	"term-deposit-calculator/internal/logger"

	"github.com/rs/zerolog"
)

func RunApp(cfg config.Config, log zerolog.Logger) error {

	return nil
}

func main() {
	cfg := config.ParseConfig()
	log := logger.ConfigureLogger(cfg)

	if err := RunApp(cfg, log); err != nil {
		log.Error().Err(err).Msg("Application failed")
		os.Exit(1)
	}
}
