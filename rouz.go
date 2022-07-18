package base

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	RouzId string
)

type Rouz struct {
	Id          string       `json:"id"`
	CreadoEn    time.Time    `json:"creacion"`
	Hilo        Hilo         `json:"hilo"`
	Comentarios []Comentario `json:"comentarios,omitempty"`
}

func New(id string) *Rouz {
	return ScanRouz(id)
}

// ScanRouz lee un arreglo de bytes y lo mapea a un struct de Rouz
func ScanRouz(id string) *Rouz {
	var err error
	bodyData := GetRouz(id)

	r := new(Rouz)

	err = json.Unmarshal(bodyData, r)
	Handle(err, fmt.Sprintf("No puedo escanearlo mi amol"))

	log.Println("Rouz escaneado satisfactoriamente...")
	// readFiles(r)
	return r
}

// GetRouz hace la petición y regresa el body de la respuesta
func GetRouz(id string) []byte {
	req, err := http.NewRequest("GET", BaseUrl+id, nil)
	Handle(err, fmt.Sprintf("¡Chocamo!\n"))

	req.AddCookie(&AccessCookie)

	response, err := Client.Do(req)
	Handle(err, fmt.Sprintf("No puedo conectarme mi amol"))
	defer response.Body.Close()

	// Mordekai pégate un tiro, con tu jodita de las rouzcoins hiciste todas las respuestas más pesadas
	const maxCapacity = 512 * 2048
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(response.Body)
	scanner.Buffer(buf, maxCapacity)

	var bodyData []byte
	for scanner.Scan() {
		line := scanner.Bytes()
		if bytes.Contains(line, []byte("window.data")) {
			_, bodyData, _ = bytes.Cut(line, []byte("   window.data = "))
			break
		}
	}
	return bodyData
}
