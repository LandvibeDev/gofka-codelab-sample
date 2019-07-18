package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID    string `json:"id" form:"id" query:"id"`
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

// e.POST("/users", SaveUser)
func (h *Handler) SaveUser(c echo.Context) error {
	ctx := c.Request().Context()

	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	insertResult, err := h.userCollection.InsertOne(ctx, u)
	if err != nil {
		return err
	}

	fmt.Println("Inserted a single documents: ", insertResult.InsertedID)
	return c.JSON(http.StatusCreated, u)
}

// e.GET("/users/:id", GetUser)
func (h *Handler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	// User ID from path `users/:id`
	id := c.Param("id")

	var user User
	if err := h.userCollection.FindOne(ctx, bson.D{{"id", id}}).Decode(&user); err != nil {
		fmt.Println(id+" not exist: ", err)
		return c.String(http.StatusNotFound, id+" not exist")
	}

	return c.JSON(http.StatusOK, user)
}

// e.PUT("/users/:id", UpdateUser)
func (h *Handler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()
	// User ID from path `users/:id`
	id := c.Param("id")

	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	updatedResult, err := h.userCollection.ReplaceOne(ctx, bson.D{{"id", id}}, u)
	if err != nil || updatedResult == nil {
		return err
	}

	if updatedResult.MatchedCount == 0 {
		return c.String(http.StatusNotFound, id+" not exist")
	}

	fmt.Println("Updated a single documents: ", updatedResult.UpsertedID)
	return c.JSON(http.StatusOK, u)
}

// e.DELETE("/users/:id", DeleteUser)
func (h *Handler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()

	// User ID from path `users/:id`
	id := c.Param("id")

	deletedResult, err := h.userCollection.DeleteOne(ctx, bson.D{{"id", id}})
	if err != nil {
		return err
	}

	fmt.Println("Deleted documents count: ", deletedResult.DeletedCount)
	return c.String(http.StatusNoContent, id+" is deleted")
}
