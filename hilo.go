package base

type Hilo struct {
	Id     string `json:"id"`
	Titulo string `json:"titulo"`
	Media  Media  `json:"media"`
}
