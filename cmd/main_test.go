package main

import (
	"bytes"
	"term-deposit-calculator/internal/config"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestRunApp_Success(t *testing.T) {
	cfg := config.Config{
		StartDeposit:   1000,
		InterestRate:   5,
		InvestmentTerm: 2,
		InterestPaid:   "at-maturity",
	}

	var logBuffer bytes.Buffer
	log := zerolog.New(&logBuffer)

	err := RunApp(cfg, log)
	assert.NoError(t, err)

	output := logBuffer.String()
	assert.Contains(t, output, "Final balance: $1100")
	assert.Contains(t, output, "interest paid at-maturity")
}

func TestRunApp_InvalidInterestPaid(t *testing.T) {
	cfg := config.Config{
		StartDeposit:   1000,
		InterestRate:   5,
		InvestmentTerm: 1,
		InterestPaid:   "invalid",
	}

	var logBuffer bytes.Buffer
	log := zerolog.New(&logBuffer)

	err := RunApp(cfg, log)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported interest payment frequency")
}
