package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() //	instanciando la libreria echo

	//	e.GET("/", holaMundo)

	//	CORS
	//	DESMARCAR ESTO PARA PROBAR CORS
	//	e.Use(middleware.CORS()) //	esto permite a cualquier dominio realizar consultas (esto solo es para pruebas)
	//	IMPORTANTE: se debe especifica que dominios puede realizar consultas

	//	como probarlo
	//	en el console de la web
	//	fetch('http://localhost:8080/api/v1/users').the(r => r.json()).then(data => console.log(data));

	startRoutes(e) //	llamando a la funcion startRoutes de route.go

	err := e.Start(":8080") //	localhost:8080

	if err != nil {
		fmt.Printf("No se pudo subir el server %v", err)
	}
}
