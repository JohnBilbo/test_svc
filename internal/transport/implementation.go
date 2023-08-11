package transport

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler interface {
	Live(c echo.Context) error
	Ready(c echo.Context) error
}

func (h *handler) Live(c echo.Context) error {
	return c.String(http.StatusOK, "live")
}

func (h *handler) Ready(c echo.Context) error {
	return c.String(http.StatusOK, "ready")
}
