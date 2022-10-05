package zapato

var storage Storage

// IMPORTANTE inicializar variable
func init() {
	storage = make(map[string]*Model)
}

// creando estructura
type Model struct {
	Marca  string `json:"marca"` //	sign: que cuando se convierta a json quitara el mayus
	Precio int    `json:"precio"`
	Color  string `json:"color"`
}

// Mapa: structura de clave y valor
type Storage map[string]*Model

// MÉTODOS
func (s Storage) Create(m *Model) *Model {
	s[m.Marca] = m    //	Marca: key | m: valores
	return s[m.Marca] //	devuelve el registro creado
}

func (s Storage) GetAll() Storage { //	devuelve el mapa
	return s
}

func (s Storage) GetByMarca(m string) *Model { //	devuelve el modelo
	if v, ok := s[m]; ok {
		return v
	}

	return nil
}

func (s Storage) Update(m string, z *Model) {
	s[m] = z //	modifica por el key
}

func (s Storage) Delete(m string) {
	delete(s, m)
}
