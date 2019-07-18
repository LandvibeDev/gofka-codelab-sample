package main

import (
	"context"

	"github.com/LandvibeDev/gofka-codelab-sample/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		e.Logger.Fatal(err)
	}

	collection := client.Database("gofka").Collection("users")

	v1 := e.Group("/api/v1")
	h := router.NewHandler(collection)
	h.Register(v1)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
