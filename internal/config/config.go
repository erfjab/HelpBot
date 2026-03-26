package config

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Debug bool `mapstructure:"DEBUG"`

	TelegramToken       string  `mapstructure:"TELEGRAM_TOKEN"`
	TelegramAdminsIDRaw string  `mapstructure:"TELEGRAM_ADMINS_ID"`
	TelegramAdminsID    []int64 `mapstructure:"-"`
}


var Cfg *Config

func LoadConfig() (*Config, error) {
	var config Config

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		var configNotFound viper.ConfigFileNotFoundError
		if !errors.As(err, &configNotFound) {
			return &config, err
		}
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return &config, err
	}

	config.TelegramAdminsID, err = parseAdminIDs(config.TelegramAdminsIDRaw)
	if err != nil {
		return &config, fmt.Errorf("TELEGRAM_ADMINS_ID: %w", err)
	}

	if err = config.validate(); err != nil {
		return &config, err
	}

	Cfg = &config
	return &config, nil
}

func (c *Config) validate() error {
	if c.TelegramToken == "" {
		return errors.New("TELEGRAM_TOKEN is required")
	}

	return nil
}

func parseAdminIDs(raw string) ([]int64, error) {
	if raw == "" {
		return []int64{}, nil
	}
	parts := strings.Split(raw, ",")
	ids := make([]int64, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		id, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid admin ID %q: %w", p, err)
		}
		ids = append(ids, id)
	}
	return ids, nil
}