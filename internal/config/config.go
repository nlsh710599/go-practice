package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port int `default:"8080" split_words:"true"`

	RpcUrl   string `default:"https://data-seed-prebsc-2-s3.binance.org:8545/" split_words:"true"`
	WsRpcUrl string `default:"wss://testnet-dex.binance.org/api/" split_words:"true"`

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
