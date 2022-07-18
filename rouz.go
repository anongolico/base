package base

import "time"

type Rouz struct {
	Id          string       `json:"id"`
	CreadoEn    time.Time    `json:"creacion"`
	Hilo        Hilo         `json:"hilo"`
	Comentarios []Comentario `json:"comentarios,omitempty"`
}
