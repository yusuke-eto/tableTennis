package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HandleRoutes() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello_world", helloWorldHandler)
	e.GET("/user", userHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
