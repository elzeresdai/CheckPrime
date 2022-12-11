package main

import (
	"check_primes/internal"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/", internal.CheckPrimes)

	// Start server
	e.Logger.Fatal(e.Start(":8070"))
}
