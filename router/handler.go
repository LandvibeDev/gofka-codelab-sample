package router

import (
	"github.com/LandvibeDev/gofka-codelab-sample/service"
)

type Handler struct {
	userService service.UserServiceInterface
	logService  service.LogServiceInterface
}

func NewHandler(u service.UserServiceInterface, l service.LogServiceInterface) *Handler {
	return &Handler{
		userService: u,
		logService:  l,
	}
}
