package respuesta

//	Estrtuctura para la nevegacion de link
type Navegacion struct {
	Descripcion string `json:"descripcion"` //	describe el tipo de recurso (puede ser la misma consulta o donde encontrar a todos los usuarios)
	Link        string `json:"link"`        //	link uri
}
