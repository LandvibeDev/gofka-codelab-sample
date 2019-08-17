package main

import (
	"fmt"
	"github.com/LandvibeDev/gofka-codelab-sample/config"
	"github.com/LandvibeDev/gofka-codelab-sample/db"
	"github.com/LandvibeDev/gofka-codelab-sample/kafka"
	"github.com/LandvibeDev/gofka-codelab-sample/router"
	"github.com/LandvibeDev/gofka-codelab-sample/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"runtime"
)

func main() {
	// Use multi core
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Load Configuration
	gofkaConfig := config.LoadConfiguration(e.Logger)

	// Connect DB
	client, err := db.New(gofkaConfig.MongoDb)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Connect Kafka
	_, err = kafka.EnsureTopic(gofkaConfig.Kafka)
	if err != nil {
		e.Logger.Fatal(err)
	}

	producer, err := kafka.GetProducer(gofkaConfig.Kafka)
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

	// Create Consumer
	consumer, err := kafka.NewConsumerConnector(gofkaConfig.Kafka)
	if err != nil {
		e.Logger.Fatal(err)
	}

	go consumer.StartPeek()

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
