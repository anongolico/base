package base

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	BaseUrl        = "https://rouzed.one/Hilo/"
	MediaURL       = "https://rouzed.one/Media/"
	AuthCookieFile = ".accessCookie"
)

var (
	client = http.Client{}
	rouzId string
)

// authenticateAnon sets the cookie used for the requests
func authenticateAnon(accessCookieValue string) {
	SetAccessCookieValue(accessCookieValue)
}

// getRouz hace la petición y regresa el body de la respuesta
func getRouz() []byte {
	req, err := http.NewRequest("GET", BaseUrl+rouzId, nil)
	handle(err, fmt.Sprintf("¡Chocamo!\n"))

	req.AddCookie(&AccessCookie)

	response, err := client.Do(req)
	handle(err, fmt.Sprintf("No puedo conectarme mi amol"))
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

// scanRouz lee un arreglo de bytes y lo mapea a un struct de Rouz
func scanRouz() *Rouz {
	var err error
	bodyData := getRouz()

	rouz := new(Rouz)

	err = json.Unmarshal(bodyData, rouz)
	handle(err, fmt.Sprintf("No puedo escanearlo mi amol"))

	log.Println("Rouz escaneado satisfactoriamente...")
	// readFiles(rouz)
	return rouz
}

// TODO: descargar archivos basado en las llaves del mapa 'files'
/*
func readFiles(rouz *Rouz) {
	for _, comentario := range rouz.Comentarios {
		if comentario.Media.Url != "" {
			_, formato, _ := strings.Cut(comentario.Media.Url, ".")
			files[formato] = append(files[formato], comentario.Media.Url)
		}
	}
}*/

// readAuthFile lee la cookie de acceso desde un archivo
func readAuthFile() {
	_, err := os.Stat(AuthCookieFile)
	if os.IsNotExist(err) {
		createAuthFile()
	}
	content, err := os.ReadFile(AuthCookieFile)
	handle(err, "error al leer archivo de credenciales")

	if string(content) == "" {
		fmt.Println("this")
	}
	authenticateAnon(string(content))
}

// createAuthFile reads the cookie from stdin and saves it to a new .accessCookie file
func createAuthFile() {
	f, err := os.Create(AuthCookieFile)
	handle(err, "error al crear archivo de credenciales")
	defer f.Close()

	fmt.Println("Pega el valor de la cookie '.AspNetCore.Identity.Application' (Ctrl+Mayus+V):")
	var s string
	_, err = fmt.Scan(&s)
	handle(err, "error al leer entrada de usuario")
	_, err = f.Write([]byte(s))
	handle(err, "error al escribir contenido de archivo")
}

func downloadFiles(r *Rouz) {
	folderName := fmt.Sprintf("%s (%s)", r.Hilo.Titulo, r.Hilo.Id)
	createRouzFolder(folderName)
	err := os.Chdir(folderName)
	handle(err, "")

	err = downloadFile(r.Hilo.Media.Url)
	for _, v := range r.Comentarios {
		if v.Media.Url != "" && !strings.Contains(v.Media.Url, "https://") {
			err = downloadFile(v.Media.Url)
			if err != nil {
				log.Printf("error al descargar archivo: %s", v.Media.Url)
			}
			// comentariosConMedia--
			// fmt.Printf("%d archivos restantes.\n", comentariosConMedia)
		}
	}

	err = os.Chdir("..")
	handle(err, "")
}

func main() {
	readAuthFile()

	fmt.Println("Pega el rouz id (Ctrl+Mayus+V):")
	_, err := fmt.Scan(&rouzId)
	handle(err, "")

	r := scanRouz()

	downloadFiles(r)

	log.Printf("Operación completada. ¡Hasta la prótsima!\n")
}

// createRouzFolder creates a new folder to store media
func createRouzFolder(name string) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		err = os.Mkdir(name, 0755)
		handle(err, "")
	}
}

func downloadFile(url string) error {
	// Get the data
	resp, err := http.Get(MediaURL + url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Create the file
	out, err := os.Create(url)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

/*
comentarios := 0
	comentariosConMedia := 0
	for _, v := range r.Comentarios {
		comentarios++
		if v.Media.Url != "" {
			comentariosConMedia++
		}
	}

	log.Printf("%d comentarios encontrados, %d con archivo adjunto", comentarios, comentariosConMedia)
*/
