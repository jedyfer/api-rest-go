package main

import (
	"api-go/usuario"
	"api-go/zapato"

	"github.com/labstack/echo/v4"
)

func startRoutes(e *echo.Echo) {
	e.POST("/zapatos", zapato.Create, usuario.ValidateJWT)
	e.POST("/login", usuario.Login)

	//	usuarios
	e.POST("/api/v1/users", usuario.Create)
	e.GET("/api/v1/users", usuario.GetAll)
	e.GET("/api/v1/users/:email", usuario.GetByEmail)

	//	params
	e.GET("/api/v1/users-paginate", usuario.GetAllPaginate)
}
