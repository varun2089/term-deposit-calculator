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
	td := calculator.TermDeposit{
		StartDeposit:  cfg.StartDeposit,
		InterestRate:  cfg.InterestRate,
		InvestmentTerm: cfg.InvestmentTerm,
		InterestPaid:  cfg.InterestPaid,
	}

	finalBalance := td.CalculateFinalBalance()

	fmt.Printf("Final balance: $%.2f\n", finalBalance)

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
