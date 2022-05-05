package routing

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// set routing
	SetPostRouting(e)

	// start server
	e.Logger.Fatal(e.Start(":8080"))
}