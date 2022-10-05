package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() //	instanciando la libreria echo

	//	e.GET("/", holaMundo)

	startRoutes(e) //	llamando a la funcion startRoutes de route.go

	err := e.Start(":8080") //	localhost:8080

	if err != nil {
		fmt.Printf("No se pudo subir el server %v", err)
	}
}
