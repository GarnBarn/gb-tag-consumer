package main

import (
	"fmt"
	"github.com/GarnBarn/common-go/database"
	"github.com/GarnBarn/common-go/httpserver"
	"github.com/GarnBarn/common-go/rabbitmq"
	"github.com/GarnBarn/gb-tag-consumer/cmd/gb-tag-delete-consumer/processor"
	"github.com/GarnBarn/gb-tag-consumer/pkg/config"
	"github.com/GarnBarn/gb-tag-consumer/pkg/repository"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var (
	appConfig config.Config
)

func init() {
	appConfig = config.Load()
}

func main() {

	rabbitMQ, err := rabbitmq.NewRabbitMQ(appConfig.RABBITMQ_CONNECTION)
	if err != nil {
		logrus.Fatal("Connect RabbitMQ Error: ", err)
	}
	// Connect Database
	db, err := database.Conn(appConfig.MYSQL_CONNECTION_STRING)
	if err != nil {
		logrus.Fatalln("Can't connect to database: ", err)
		return
	}

	// Start HealthChecking Server
	go func() {
		httpServer := httpserver.NewHttpServer()
		logrus.Info("Listening and serving HTTP on :", appConfig.DELETE_TAG_HTTP_SERVER_PORT)
		httpServer.Run(fmt.Sprint(":", appConfig.DELETE_TAG_HTTP_SERVER_PORT))
	}()

	// Create Repository
	tagRepository := repository.NewTagRepository(db)

	// Create Processor
	processor := processor.NewProcessor(rabbitMQ.GetPublisher(), tagRepository)

	consumer, err := rabbitMQ.Consume(processor, rabbitmq.ConsumerConfig{
		MaxRetry:           appConfig.RABBITMQ_MAXIMUM_RETRY,
		FailoverExchange:   appConfig.RABBITMQ_TAG_EXCHANGE,
		DeadLetterExchange: appConfig.RABBITMQ_DEAD_LETTERING_EXCHANGE,
		ConsumeQueue:       appConfig.RABBITMQ_TAG_DELETE_QUEUE,
	})

	if err != nil {
		logrus.Fatal(err)
	}

	if err != nil {
		logrus.Fatal(err)
	}

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	<-gracefulStop

	logrus.Info("Gracefully shutting down.")
	consumer.Close()
	rabbitMQ.CloseConnection()

	logrus.Info("Successfully shutting down the amqp. Bye!!")
}
