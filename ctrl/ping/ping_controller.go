package ping

import (
	"github.com/labstack/echo"
	"net/http"
)

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}