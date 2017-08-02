package main

import (
	"os"

	"github.com/spf13/viper"
)

// Constants for Environment variables
const (
	Environment    = "SKYNET__ENVIRONMENT"
	CoinbaseKey    = "SKYNET__COINBASE_KEY"
	CoinbaseSecret = "SKYNET__COINBASE_SECRET"
	CoinbasePhrase = "SKYNET__COINBASE_PHRASE"
	PidPath        = "SKYNET__PIDPATH"
	PidFile        = "SKYNET__PIDFILE"
)

// ConfigReader represents configuration reader
type ConfigReader interface {
	Get(string) interface{}
	GetString(string) string
	GetInt(string) int
	GetBool(string) bool
	GetStringMap(string) map[string]interface{}
	GetStringMapString(string) map[string]string
	GetStringSlice(string) []string
	SetDefault(string, interface{})
}

// DefaultSettings is the function for configuring defaults
type DefaultSettings func(config ConfigReader)

// ConfigDefaults - returns the defauls of the config passed
func ConfigDefaults(config ConfigReader) {
	Defaults(config)
}

// Defaults is the default settings functor
func Defaults(config ConfigReader) {
	config.SetDefault("environment", GetEnv(Environment, "local"))
	config.SetDefault("coinbase_key", GetEnv(CoinbaseKey, ""))
	config.SetDefault("coinbase_secret", GetEnv(CoinbaseSecret, ""))
	config.SetDefault("coinbase_phrase", GetEnv(CoinbasePhrase, ""))
	config.SetDefault("pidfile", GetEnv(PidFile, "skynet.pid"))
	config.SetDefault("pidpath", GetEnv(PidPath, "/var/run/skynet"))
}

// GetEnv - pull values or set defaults.
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	return value
}

// LoadConfig - returns configuration for a particular app
func LoadConfig(defaultSetup DefaultSettings) ConfigReader {
	config := viper.New()

	Defaults(config)

	return config
}
