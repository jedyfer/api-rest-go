package usuario

import (
	"api-go/respuesta"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	u := &Model{}

	err := c.Bind(u)

	if err != nil {
		r := respuesta.Model{
			MensajeError: respuesta.MensajeError{
				Codigo:    "E003",
				Contenido: "El objeto ha sido enviado mal",
			},
		}

		return c.JSON(http.StatusBadRequest, r)
	}

	d := storage.Login(u.Email, u.Password)

	if d == nil {
		r := respuesta.Model{
			MensajeError: respuesta.MensajeError{
				Codigo:    "L001",
				Contenido: "Usuario o password incorrecto",
			},
		}

		return c.JSON(http.StatusBadRequest, r)
	}

	d.Password = "" //	para no enviar el password
	r := respuesta.Model{
		MensajeOK: respuesta.MensajeOK{
			Codigo:    "O001",
			Contenido: "Logeado",
		},
		Data: d,
	}

	return c.JSON(http.StatusOK, r)
}
