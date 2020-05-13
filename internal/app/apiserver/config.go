package apiserver

import "github.com/RomanAvdeenko/http-rest-api/internal/app/store"

// Config ...
type Config struct {
	BindAddr uint   `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: 8080,
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
