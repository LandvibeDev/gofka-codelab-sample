package main

import (
	"github.com/LandvibeDev/gofka-codelab-sample/db"
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

	client, err := db.New()
	if err != nil {
		e.Logger.Fatal(err)
	}

	userService := service.NewUserService(client)

	v1 := e.Group("/api/v1")
	h := router.NewHandler(userService)
	h.Register(v1)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
