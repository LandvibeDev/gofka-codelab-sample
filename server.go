package main

import (
	"github.com/LandvibeDev/gofka-codelab-sample/db"
	"github.com/LandvibeDev/gofka-codelab-sample/kafka"
	"github.com/LandvibeDev/gofka-codelab-sample/router"
	"github.com/LandvibeDev/gofka-codelab-sample/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect DB
	client, err := db.New()
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Connect Kafka
	kafkaConfig := kafka.KafkaConfig{Hosts: "172.17.0.1:9093"}
	topicConfig := kafka.KafkaTopicConfig{Topic: service.LogTopic, NumPartitions: 1, ReplicationFactor: 1}
	_, err = kafka.EnsureTopic(topicConfig, kafkaConfig)
	if err != nil {
		e.Logger.Fatal(err)
	}

	producer, err := kafka.GetProducer(kafkaConfig)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer producer.Close()

	// Create Service
	userService := service.NewUserService(client)
	logService := service.NewLogService(producer)

	// Create Router
	v1 := e.Group("/api/v1")
	h := router.NewHandler(userService, logService)
	h.Register(v1)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
