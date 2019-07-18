package router

import "github.com/labstack/echo"

func (h *Handler) Register(v1 *echo.Group) {
	user := v1.Group("/users")
	user.POST("", h.SaveUser)
	user.GET("/:id", h.GetUser)
	user.PUT("/:id", h.UpdateUser)
	user.DELETE("/:id", h.DeleteUser)
}
