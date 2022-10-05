package main

import (
	"api-go/zapato"

	"github.com/labstack/echo/v4"
)

func startRoutes(e *echo.Echo) {
	e.POST("/zapatos", zapato.Create)
}
