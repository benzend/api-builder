package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func User(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}