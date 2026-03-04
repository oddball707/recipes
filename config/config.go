package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	DB   DBConfig
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

var cfg *Config

// Init initializes the configuration from environment and files
func Init() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	v := viper.New()

	// Detect environment
	env := detectEnvironment()

	// Set defaults
	v.SetDefault("port", "8080")

	// Load from config file for local environment
	if env == "local" {
		v.SetConfigName("config.local")
		v.SetConfigType("json")
		v.AddConfigPath(".")

		if err := v.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return nil, fmt.Errorf("error reading config file: %w", err)
			}
			// File not found is okay, we'll use env vars
		}
	}

	// Bind environment variables
	v.SetEnvPrefix("")
	v.BindEnv("db_host", "DB_HOST")
	v.BindEnv("db_port", "DB_PORT")
	v.BindEnv("db_user", "DB_USER")
	v.BindEnv("db_password", "DB_PASSWORD")
	v.BindEnv("db_name", "DB_NAME")
	v.BindEnv("port", "PORT")

	// Create config struct
	cfg = &Config{
		DB: DBConfig{
			Host:     v.GetString("db_host"),
			Port:     v.GetString("db_port"),
			User:     v.GetString("db_user"),
			Password: v.GetString("db_password"),
			Name:     v.GetString("db_name"),
		},
		Port: v.GetString("port"),
	}

	// Validate configuration
	if err := validate(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// Get returns the initialized config
func Get() *Config {
	if cfg == nil {
		panic("config not initialized - call Init() first")
	}
	return cfg
}

// detectEnvironment determines the runtime environment
func detectEnvironment() string {
	// Check if running in AWS Lambda
	if _, ok := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME"); ok {
		return "lambda"
	}

	// Check if running in Docker
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return "docker"
	}

	// Default to local
	return "local"
}

// validate ensures all required configuration values are present
func validate(cfg *Config) error {
	if cfg.DB.Host == "" {
		return fmt.Errorf("DB_HOST is required")
	}
	if cfg.DB.Port == "" {
		return fmt.Errorf("DB_PORT is required")
	}
	if cfg.DB.User == "" {
		return fmt.Errorf("DB_USER is required")
	}
	if cfg.DB.Password == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}
	if cfg.DB.Name == "" {
		return fmt.Errorf("DB_NAME is required")
	}
	return nil
}
