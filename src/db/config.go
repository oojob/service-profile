package db

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config database config
type Config struct {
	DatabaseURI   string
	RedisURI      string
	AccessSecret  string
	RefreshSecret string
}

// InitConfig initialize configuration
func InitConfig() (*Config, error) {
	config := &Config{
		DatabaseURI:   viper.GetString("databaseuri"),
		RedisURI:      viper.GetString("redisuri"),
		AccessSecret:  viper.GetString("accesssecret"),
		RefreshSecret: viper.GetString("refreshsecret"),
	}

	if config.DatabaseURI == "" {
		return nil, fmt.Errorf("DatabaseURI must be set")
	}
	if config.RedisURI == "" {
		return nil, fmt.Errorf("RedisURI must be set")
	}
	if config.AccessSecret == "" {
		return nil, fmt.Errorf("AccessSecret must be set")
	}
	if config.RefreshSecret == "" {
		return nil, fmt.Errorf("RefreshSecret must be set")
	}

	return config, nil
}
