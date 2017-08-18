package main

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

// Constants for Environment variables
const (
	Name = "skynet"
	Env  = "env"

	Environment    = "environment"
	EnvironmentVar = "SKYNET__ENVIRONMENT"

	CoinbaseKey    = "coinbase_key"
	CoinbaseKeyVar = "SKYNET__COINBASE_KEY"

	CoinbaseSecret    = "coinbase_secret"
	CoinbaseSecretVar = "SKYNET__COINBASE_SECRET"

	CoinbasePhrase    = "coinbase_phrase"
	CoinbasePhraseVar = "SKYNET__COINBASE_PHRASE"

	PidFull = "pidfull"

	PidFile    = "pidfile"
	PidFileVar = "SKYNET__PIDFILE"

	PidPath    = "pidpath"
	PidPathVar = "SKYNET__PIDPATH"

	Local      = "local"
	Production = "production"
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
	config.SetDefault(Environment, GetEnv(EnvironmentVar, Local))
	config.SetDefault(CoinbaseKey, GetEnv(CoinbaseKeyVar, ""))
	config.SetDefault(CoinbaseSecret, GetEnv(CoinbaseSecretVar, ""))
	config.SetDefault(CoinbasePhrase, GetEnv(CoinbasePhraseVar, ""))
	config.SetDefault(PidFile, GetEnv(PidFileVar, "skynet.pid"))
	config.SetDefault(PidPath, GetEnv(PidPathVar, "/var/run/skynet"))
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
func LoadConfig(defaultSetup DefaultSettings) (ConfigReader, error) {
	config := viper.New()

	Defaults(config)

	var e error

	if config.GetString(CoinbaseSecret) == "" || config.GetString(CoinbaseKey) == "" || config.GetString(CoinbasePhrase) == "" {
		e = errors.New("missing or invalid CoinbaseSecret/CoinbaseKey/CoinbasePhrase")
	} else {
		e = nil
	}

	return config, e
}
