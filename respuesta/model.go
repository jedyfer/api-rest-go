package respuesta

//	estructura de respuesta
type Model struct {
	MensajeOK    MensajeOK    `json:"mensaje_ok"`
	MensajeError MensajeError `json:"mensaje_error"`

	Data interface{} `json:"data"` //	permite devolver cualquier tipo de estrucutra para procesarla
}

type MensajeOK struct {
	Codigo    string `json:"codigo"`
	Contenido string `json:"contenido"`
}

type MensajeError struct {
	Codigo    string `json:"codigo"`
	Contenido string `json:"contenido"`
}
