package base

import "time"

type Comentario struct {
	Id        string    `json:"id"`
	CreadoEn  time.Time `json:"creacion"`
	Contenido string    `json:"contenido"`
	Media     Media     `json:"media"`
}
