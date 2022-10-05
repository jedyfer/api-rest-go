package zapato

import (
	"api-go/respuesta"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	m := &Model{} //	almacenando la data
	err := c.Bind(m)

	if err != nil {
		r := respuesta.Model{
			MensajeError: respuesta.MensajeError{
				Codigo:    "E102", //	codigo de la empresa
				Contenido: "El objeto zapato está mal enviado",
			},
		}

		return c.JSON(http.StatusBadRequest, r)
	}

	d := storage.Create(m)

	r := respuesta.Model{
		MensajeOK: respuesta.MensajeOK{
			Codigo:    "A001",
			Contenido: "Zapato creado crrectamente",
		},
		Data: d,
	}

	return c.JSON(http.StatusCreated, r) //	201: statusCreated
}
