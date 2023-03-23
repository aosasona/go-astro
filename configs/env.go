package configs

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type AppEnv string

const (
	PRODUCTION  AppEnv = "production"
	DEVELOPMENT AppEnv = "development"
	STAGING     AppEnv = "staging"
)

type Config struct {
	AppName        string `mapstructure:"APP_NAME"`
	AppEnv         AppEnv `mapstructure:"APP_ENV"`
	Port           string `mapstructure:"PORT"`
	AllowedOrigins string `mapstructure:"ALLOWED_ORIGINS"`

	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`

	RedisURL string `mapstructure:"REDIS_URL"`
}

func LoadEnv(path string) (*Config, error) {
	c := new(Config)

	err := c.LoadWithViper(path)
	if err != nil {
		return c, err
	}

	return c.LoadDefaults(), nil
}

func (c *Config) LoadDefaults() *Config {
	if c.AppName == "" {
		c.AppName = "Cool App"
	}

	if c.Port == "" {
		log.Warn("Invalid port detected, defaulting to `8080`")
		c.Port = "8080"
	}

	if c.AppEnv == "" ||
		(c.AppEnv != PRODUCTION && c.AppEnv != DEVELOPMENT && c.AppEnv != STAGING) {
		log.Warn("Invalid app environment detected, defaulting to `development`")
		c.AppEnv = DEVELOPMENT
	}

	if c.AllowedOrigins == "" {
		c.AllowedOrigins = "*"
	}

	return c
}

func (c *Config) LoadWithViper(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.BindEnv("PORT", "PORT")
	viper.BindEnv("APP_NAME", "APP_NAME")
	viper.BindEnv("APP_ENV", "APP_ENV")
	viper.BindEnv("ALLOWED_ORIGINS", "ALLOWED_ORIGINS")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("unable to load environment variables: %v", err.Error())
		}
	}

	if err := viper.Unmarshal(c); err != nil {
		return fmt.Errorf("unable to load marshal into struct: %v", err.Error())
	}

	return nil
}
