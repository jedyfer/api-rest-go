package usuario

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

var storage Storage

func init() {
	storage = make(map[string]*Model)

	u1 := &Model{
		FirstName: "Jediael",
		Email:     "jedyfer@gmail.com",
		Password:  "123",
	}

	u2 := &Model{
		FirstName: "Juan",
		Email:     "juan@gmail.com",
		Password:  "123",
	}

	storage.Create(u1)
	storage.Create(u2)
}

type Model struct {
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Storage map[string]*Model

// MÉTODOS
func (s Storage) Create(m *Model) *Model {
	s[m.Email] = m    //	Marca: key | m: valores
	return s[m.Email] //	devuelve el registro creado
}

func (s Storage) GetAll() Storage { //	devuelve el mapa
	return s
}

func (s Storage) GetByMarca(e string) *Model { //	devuelve el modelo
	if v, ok := s[e]; ok {
		return v
	}

	return nil
}

func (s Storage) GetByEmail(e string) *Model {
	if v, ok := s[e]; ok {
		return v
	}

	return nil
}

func (s Storage) Update(e string, z *Model) *Model { //	update by email
	s[e] = z //	modifica por el key
	return s[e]
}

func (s Storage) Delete(e string) { //	delete by email
	delete(s, e)
}

func (s Storage) Login(e, p string) *Model {
	for _, v := range s {
		if v.Email == e && v.Password == p {
			return v
		}
	}

	return nil
}

// CON PARAMETROS de paginacion
func (s Storage) GetAllPaginate(l, p int) []*Model {
	us := make([]*Model, 0, len(s))

	for _, v := range s {
		us = append(us, v)
		fmt.Println(v)
	}

	fmt.Println(us)

	offset := l*p - l
	r := us[offset : l*p]

	return r
}

type Claim struct {
	Usuario            Model
	jwt.StandardClaims //	info sobre el tiempo de expiracion del token (fecha de exp. - quien declaro ese token)
}
