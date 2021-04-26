package config

import (
	"log"

	v "github.com/spf13/viper"
)

func setDefaults() {
	v.SetDefault("PORT", "8080")
	v.SetDefault("SERVER_TYPE", "grpc")
}

func LoadConfig() {
	setDefaults()

	v.SetConfigFile(".env")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading the config file: %s", err)
	}
}
