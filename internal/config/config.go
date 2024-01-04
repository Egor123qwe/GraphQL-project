package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	GrpcPort int
}

func New() *Config {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	return &Config{
		GrpcPort: viper.GetInt("grpc_server_port"),
	}
}
