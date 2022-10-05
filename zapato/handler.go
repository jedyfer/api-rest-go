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
				"E102", //	codigo de la empresa
				"El objeto zapato está mal enviado",
			},
		}

		return c.JSON(http.StatusBadRequest, r)
	}

	d := storage.Create(m)

	r := respuesta.Model{
		MensajeOK: respuesta.MensajeOK{
			"A001",
			"Zapato creado crrectamente",
		},
		Data: d,
	}

	return c.JSON(http.StatusCreated, r) //	201: statusCreated
}
