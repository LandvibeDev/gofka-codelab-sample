package router

import (
	"fmt"
	"net/http"

	"github.com/LandvibeDev/gofka-codelab-sample/model"
	"github.com/labstack/echo"
)

// e.POST("/users", SaveUser)
func (h *Handler) SaveUser(c echo.Context) error {
	ctx := c.Request().Context()

	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	user, err := h.userService.Create(&ctx, u)
	if err != nil {
		return err
	}

	fmt.Println("Inserted a single documents: ", user)
	return c.JSON(http.StatusCreated, u)
}

// e.GET("/users/:id", GetUser)
func (h *Handler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	// User ID from path `users/:id`
	id := c.Param("id")

	user, err := h.userService.GetByID(&ctx, id)
	if err != nil {
		return c.String(http.StatusNotFound, id+" not exist")

	}
	return c.JSON(http.StatusOK, user)
}

// e.PUT("/users/:id", UpdateUser)
func (h *Handler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()
	// User ID from path `users/:id`
	id := c.Param("id")

	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if _, err := h.userService.GetByID(&ctx, id); err != nil {
		return c.String(http.StatusNotFound, id+" not exist")
	}

	user, err := h.userService.Update(&ctx, u)

	if err != nil {
		return err
	}

	fmt.Println("Updated a single documents: ", user)
	return c.JSON(http.StatusOK, u)
}

// e.DELETE("/users/:id", DeleteUser)
func (h *Handler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()

	// User ID from path `users/:id`
	id := c.Param("id")

	if _, err := h.userService.GetByID(&ctx, id); err != nil {
		return c.String(http.StatusNotFound, id+" not exist")
	}

	if err := h.userService.Delete(&ctx, id); err != nil {
		return err
	}

	fmt.Println("Deleted document id: ", id)
	return c.String(http.StatusNoContent, id+" is deleted")
}
