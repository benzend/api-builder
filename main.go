package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	IndexRoutes "api-builder/routes/index"
	UserRoutes "api-builder/routes/users"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", IndexRoutes.Home)
	e.GET("/users", UserRoutes.Users)
	e.GET("/users/:id", UserRoutes.User)

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}