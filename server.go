package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// User
type User struct {
	ID    string `json:"id" form:"id" query:"id"`
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

var UserDB map[string]User = make(map[string]User)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// e.POST("/users", saveUser)
func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if _, ok := UserDB[u.ID]; !ok {
		UserDB[u.ID] = *u
	}

	return c.JSON(http.StatusCreated, u)
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	if user, ok := UserDB[id]; ok {
		return c.JSON(http.StatusOK, user)
	} else {
		return c.String(http.StatusNotFound, id+" not exist")
	}
}

// e.PUT("/users/:id", updateUser)
func updateUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if _, ok := UserDB[id]; ok {
		UserDB[id] = *u
	} else {
		return c.String(http.StatusNotFound, id+" not exist")
	}

	return c.JSON(http.StatusOK, u)
}

// e.DELETE("/users/:id", deleteUser)
func deleteUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	if _, ok := UserDB[id]; ok {
		delete(UserDB, id)
	} else {
		return c.String(http.StatusNotFound, id+" not exist")
	}

	return c.String(http.StatusOK, id+" is deleted")
}
