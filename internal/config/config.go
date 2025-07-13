package config

import (
	"github.com/alecthomas/kong"
)

type Config struct {
	LogLevel string `env:"LOG_LEVEL" default:"INFO" help:"Log level for the application"`
}

func ParseConfig() Config {
	var config Config
	kong.Parse(&config)
	return config
}
