package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Env                              string
	CREATE_TAG_HTTP_SERVER_PORT      string `envconfig:"CREATE_TAG_HTTP_SERVER_PORT" default:"3004"`
	DELETE_TAG_HTTP_SERVER_PORT      string `envconfig:"DELETE_TAG_HTTP_SERVER_PORT" default:"3005"`
	MYSQL_CONNECTION_STRING          string `envconfig:"MYSQL_CONNECTION_STRING"`
	RABBITMQ_CONNECTION              string `envconfig:"RABBITMQ_CONNECTION" default:"amqp://guest:guest@localhost:5672"`
	RABBITMQ_MAXIMUM_RETRY           int    `envconfig:"RABBITMQ_MAXIMUM_RETRY" default:"3"`
	RABBITMQ_DEAD_LETTERING_EXCHANGE string `envconfig:"RABBITMQ_DEAD_LETTERING_EXCHANGE" default:"tag.fail"`
	RABBITMQ_TAG_EXCHANGE            string `envconfig:"RABBITMQ_TAG_EXCHANGE" default:"tag"`
	RABBITMQ_TAG_CREATE_QUEUE        string `envconfig:"RABBITMQ_TAG_CREATE_QUEUE" default:"tag.create.update"`
	RABBITMQ_TAG_DELETE_QUEUE        string `envconfig:"RABBITMQ_TAG_DELETE_QUEUE" default:"tag.delete.update""`
}

func Load() Config {
	var config Config
	ENV, ok := os.LookupEnv("ENV")
	if !ok {
		// Default value for ENV.
		ENV = "dev"
	}
	// Load the .env file only for dev env.
	ENV_CONFIG, ok := os.LookupEnv("ENV_CONFIG")
	if !ok {
		ENV_CONFIG = "./.env"
	}

	err := godotenv.Load(ENV_CONFIG)
	if err != nil {
		logrus.Warn("Can't load env file")
	}

	envconfig.MustProcess("", &config)
	config.Env = ENV

	return config
}
