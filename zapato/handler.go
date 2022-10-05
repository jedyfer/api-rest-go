package zapato

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	m := &Model{} //	almacenando la data
	err := c.Bind(m)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "El objeto enviado no es el correcto")
	}

	storage.Create(m)
	return c.JSON(http.StatusCreated, "ok")
}
