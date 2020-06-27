package interfaces

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func userHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hoge")
}
