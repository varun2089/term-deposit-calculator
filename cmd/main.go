package main

import (
	"fmt"
	"os"
	"term-deposit-calculator/internal/calculator"
	"term-deposit-calculator/internal/config"
	"term-deposit-calculator/internal/logger"

	"github.com/rs/zerolog"
)

func RunApp(cfg config.Config, log zerolog.Logger) error {
	termDeposit := calculator.TermDeposit{
		StartDeposit:   cfg.StartDeposit,
		InterestRate:   cfg.InterestRate,
		InvestmentTerm: cfg.InvestmentTerm,
		InterestPaid:   cfg.InterestPaid,
	}

	finalBalance, err := termDeposit.CalculateFinalBalance()
	if err != nil {
		return fmt.Errorf("failed to calculate final balance: %w", err)

	}

	log.Info().Msgf("Final balance: $%d, interest paid %s", int(finalBalance), cfg.InterestPaid)
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
