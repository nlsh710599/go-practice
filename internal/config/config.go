package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port int `default:"8080" split_words:"true"`

	RpcUrl string `default:"https://nd-391-648-435.p2pify.com/899fe2afc3bc6426d419f89800c2d871" split_words:"true"`
	WsUrl  string `default:"wss://ws-nd-391-648-435.p2pify.com/899fe2afc3bc6426d419f89800c2d871" split_words:"true"`

	PostgresHost     string `default:"localhost" split_words:"true"`
	PostgresPort     int    `default:"5432" split_words:"true"`
	PostgresDatabase string `default:"postgres" split_words:"true"`
	PostgresUser     string `default:"postgres" split_words:"true"`
	PostgresPassword string `default:"docker" split_words:"true"`

	ConfirmationBlockCount int `default:"20" split_words:"true"`
}

var config Config

func Get() *Config {
	return &config
}

func init() {
	err := envconfig.Process("", &config)
	if err != nil {
		log.Panic(err.Error())
	}
}
