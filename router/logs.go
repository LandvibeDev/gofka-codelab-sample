package router

import (
	"github.com/LandvibeDev/gofka-codelab-sample/kafka/message"
	"github.com/LandvibeDev/gofka-codelab-sample/service"
	"github.com/labstack/echo"
	"net/http"
)

// e.Post("/logs", WriteLogs)
func (h *Handler) WriteLogs(c echo.Context) error {
	l := new(message.LogMessage)
	if err := c.Bind(l); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := h.logService.Send(service.LogTopic, *l)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, l)
}
