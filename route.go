package main

import (
	"api-go/usuario"
	"api-go/zapato"

	"github.com/labstack/echo/v4"
)

func startRoutes(e *echo.Echo) {
	e.POST("/zapatos", zapato.Create, usuario.ValidateJWT)
	e.POST("/login", usuario.Login)
}
