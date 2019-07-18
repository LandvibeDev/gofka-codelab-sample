package router

import "go.mongodb.org/mongo-driver/mongo"

type Handler struct {
	userCollection *mongo.Collection
}

func NewHandler(c *mongo.Collection) *Handler {
	return &Handler{
		userCollection: c,
	}
}
