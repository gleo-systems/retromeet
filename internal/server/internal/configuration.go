package server

import "github.com/gleo-systems/retromeet/internal/env"

type Config struct {
	Port     int
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Logging  bool
}

func readConfiguration() *Config {
	return &Config{
		Port: env.GetIntVar("port"),
		Database: DatabaseConfig{
			Host:     env.GetVar("host"),
			Port:     env.GetIntVar("port"),
			Database: env.GetVar("database"),
			Username: env.GetVar("username"),
			Password: env.GetVar("password"),
			Logging:  env.GetBoolVar("logging"),
		},
	}
}
