package usuario

import (
	"api-go/respuesta"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// Crear usuario
func Create(c echo.Context) error {
	u := &Model{}
	err := c.Bind(u)

	if err != nil {
		r := respuesta.Model{
			MensajeError: respuesta.MensajeError{
				Codigo:    "U001",
				Contenido: "El objeto usuario no tiene la estructura correcta",
			},
		}

		return c.JSON(http.StatusBadRequest, r)
	}

	d := storage.Create(u)

	r := respuesta.Model{
		MensajeOK: respuesta.MensajeOK{
			Codigo:    "U200",
			Contenido: "Usuario creado correctamente",
		},
		Data: d,
	}

	return c.JSON(http.StatusCreated, r)
}

func Update(c echo.Context) error {
	u := &Model{}
	email := c.Param("email")

	err := c.Bind(u)

	if err != nil {
		r := respuesta.Model{
			MensajeError: respuesta.MensajeError{
				Codigo:    "U001",
				Contenido: "El objeto usuario no tiene la estructura correcta",
			},
		}

		return c.JSON(http.StatusBadRequest, r)
	}

	d := storage.Update(email, u)

	r := respuesta.Model{
		MensajeOK: respuesta.MensajeOK{
			Codigo:    "U201",
			Contenido: "Usuario actualizado correctamente",
		},
		Data: d,
	}

	return c.JSON(http.StatusOK, r)
}

func Delete(c echo.Context) error {
	email := c.Param("email")

	storage.Delete(email)

	r := respuesta.Model{
		MensajeOK: respuesta.MensajeOK{
			Codigo:    "U202",
			Contenido: "Usuario eliminado correctamente",
		},
	}

	return c.JSON(http.StatusOK, r)
}

func GetByEmail(c echo.Context) error {
	email := c.Param("email")

	u := storage.GetByEmail(email)

	if u == nil {
		r := respuesta.Model{
			MensajeError: respuesta.MensajeError{
				Codigo:    "U003",
				Contenido: "El email no se encuentra",
			},
		}

		return c.JSON(http.StatusNotFound, r)
	}

	r := respuesta.Model{
		MensajeOK: respuesta.MensajeOK{
			Codigo:    "U204",
			Contenido: "Consultado correctamente",
		},
		Data: u,
	}

	return c.JSON(http.StatusOK, r)
}

func GetAll(c echo.Context) error {
	us := storage.GetAll()

	r := respuesta.Model{
		MensajeOK: respuesta.MensajeOK{
			Codigo:    "U205",
			Contenido: "Consultado correctamente",
		},
		Data: us,
	}

	return c.JSON(http.StatusOK, r)
}

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
	token, err := generateJWT(*d)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "no se pudo generar el token")
	}

	type logueo struct {
		Usuario Model
		Token   string
	}

	l := logueo{
		*d,
		token,
	}

	r := respuesta.Model{
		MensajeOK: respuesta.MensajeOK{
			Codigo:    "O001",
			Contenido: "Logeado",
		},
		Data: l,
	}

	return c.JSON(http.StatusOK, r)
}

// getTokenFromAuthorizationHeader busca el token del header Authorization
func getTokenFromAuthorizationHeader(r *http.Request) (string, error) {
	ah := r.Header.Get("Authorization")
	if ah == "" {
		return "", errors.New("el encabezado no contiene la autorización")
	}

	// Should be a bearer token
	if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
		return ah[7:], nil
	} else {
		return "", errors.New("el header no contiene la palabra Bearer")
	}
}

// getTokenFromURLParams busca el token de la URL
func getTokenFromURLParams(r *http.Request) (string, error) {
	ah := r.URL.Query().Get("authorization")
	if ah == "" {
		return "", errors.New("la URL no contiene la autorización")
	}

	return ah, nil
}
