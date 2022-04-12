package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Users(c echo.Context) error {
	return c.String(http.StatusOK, "Ben, Geoffry, Harry")
}