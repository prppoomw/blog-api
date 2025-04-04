package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Profile     string `mapstructure:"PROFILE"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	MongoDBHost string `mapstructure:"MONGODB_HOST"`
}

func LoadConfig() *Config {
	config := Config{}
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("../")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Not found file .env, ", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Environment can't be loaded, ", err)
	}

	if config.Profile == "dev" {
		log.Println("The App is running in development env")
	}

	return &config
}
