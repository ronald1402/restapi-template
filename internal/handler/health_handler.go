package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restapi/commons"
	"sync"
)

type HealthCheckHandler struct {
	IsUnHealthy bool
	mut         *sync.RWMutex
}

func NewHealthCheckHandler(e *echo.Echo) {
	handler := &HealthCheckHandler{
		mut: new(sync.RWMutex),
	}

	e.GET("/health", handler.Check)
}
func (h *HealthCheckHandler) Check(c echo.Context) error {
	if h.IsUnHealthy {
		return c.JSON(http.StatusServiceUnavailable, commons.Response{Message: "service unavailable"})
	}

	return c.JSON(http.StatusOK, commons.Response{
		Code:    200,
		Message: "Ok",
	})
}

func (h *HealthCheckHandler) SetUnhealthy() {
	h.mut.Lock()
	h.IsUnHealthy = true
	h.mut.Unlock()
}
