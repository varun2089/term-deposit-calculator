package config

import (
	"github.com/alecthomas/kong"
)

type Config struct {
	LogLevel      string `env:"LOG_LEVEL" default:"INFO" help:"Log level for the application"`
	StartDeposit  float64 `name:"start-deposit" help:"Start deposit amount (e.g. 10000)" required:""`
	InterestRate  float64 `name:"interest-rate" help:"Interest rate (e.g. 1.10)" required:""`
	InvestmentTerm int    `name:"investment-term" help:"Investment term in years (e.g. 3)" required:""`
	InterestPaid  string `name:"interest-paid" help:"Interest paid frequency (monthly, quarterly, annually, at-maturity)" enum:"monthly,quarterly,annually,at-maturity" required:""`
}

func ParseConfig() Config {
	var config Config
	kong.Parse(&config)
	return config
}
