package configs

import (
	"fmt"
	"reflect"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type AppEnv string

var GlobalConfig *Config // temporary

const (
	PRODUCTION  AppEnv = "production"
	DEVELOPMENT AppEnv = "development"
	STAGING     AppEnv = "staging"
)

type Config struct {
	IsDev              bool   `mapstructure:"-"`
	AppName            string `mapstructure:"APP_NAME"`
	AppEnv             AppEnv `mapstructure:"APP_ENV"`
	AppURL             AppEnv `mapstructure:"APP_URL"`
	AccessTokenSecret  string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string `mapstructure:"REFRESH_TOKEN_SECRET"`
	Port               string `mapstructure:"PORT"`
	AllowedOrigins     string `mapstructure:"ALLOWED_ORIGINS"`
	TrialDuration      uint   `mapstructure:"TRIAL_DURATION"`
	DSN                string `mapstructure:"DSN"`
	RedisDSN           string `mapstructure:"REDIS_DSN"`
}

func init() {
	var err error
	GlobalConfig, err = LoadEnv(".")
	if err != nil {
		panic("unable to read config")
	}
}

func LoadEnv(path string) (*Config, error) {
	c := new(Config)

	err := c.LoadWithViper(path)
	if err != nil {
		return c, err
	}

	c.IsDev = c.AppEnv == DEVELOPMENT
	return c.LoadDefaults(), nil
}

func (c *Config) LoadDefaults() *Config {
	if c.AppName == "" {
		c.AppName = "Sidekyk"
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

	if c.TrialDuration == 0 {
		c.TrialDuration = 1
	}

	return c
}

func (c *Config) LoadWithViper(path string) error {
	viper.AutomaticEnv()

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	reflectType := reflect.TypeOf(*c)

	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		tag := field.Tag.Get("mapstructure")
		viper.BindEnv(tag)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("unable to read from file: %v", err.Error())
		}
	}

	if err := viper.Unmarshal(c); err != nil {
		return fmt.Errorf("unable to load marshal into struct: %v", err.Error())
	}

	return nil
}
