package handler

import (
	"github.com/labstack/echo/v4"
	"restapi/internal/service"
)

type handler struct {
	svc service.Service
}

func NewHandler(e *echo.Echo, svc service.Service) {
	_ = &handler{
		svc: svc,
	}
}
