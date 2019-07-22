package router

import (
	"github.com/LandvibeDev/gofka-codelab-sample/service"
)

type Handler struct {
	userService service.UserServiceInterface
}

func NewHandler(u service.UserServiceInterface) *Handler {
	return &Handler{
		userService: u,
	}
}
