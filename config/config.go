package config

import (
	"log"

	"github.com/spf13/viper"
)

type appConfig struct {
	HttpAddr string
}

var cfg = &appConfig{}

func InitViper() {
	config := viper.New()
	config.SetConfigName("local")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	config.BindEnv("HttpAddr", "HTTP_ADDR")

	if err := config.MergeInConfig(); err != nil {
		log.Fatal(err, "Failed to read configuration")
	}

	if err := config.Unmarshal(cfg); err != nil {
		log.Fatal(err)
	}
}

func HttpAddr() string {
	return cfg.HttpAddr
}
