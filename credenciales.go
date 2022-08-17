package base

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	AccessCookie = http.Cookie{
		Name:     ".AspNetCore.Identity.Application",
		Value:    "",
		Secure:   true,
		HttpOnly: true,
	}
)

const AuthCookieFile = "accessCookie.txt"

func SetAccessCookieValue(v string) {
	AccessCookie.Value = v
}

// ReadAuthFile lee la cookie de acceso desde un archivo
func ReadAuthFile() {
	_, err := os.Stat(AuthCookieFile)
	if os.IsNotExist(err) {
		createAuthFile()
	}
	content, err := os.ReadFile(AuthCookieFile)
	Handle(err, "error al leer archivo de credenciales")

	contentString := string(content)
	contentString = strings.TrimSpace(contentString)

	// TODO: do something if file exists and it's empty

	SetAccessCookieValue(contentString)
}

// createAuthFile reads the cookie from stdin and saves it to a new accessCookie.txt file
func createAuthFile() {
	f, err := os.Create(AuthCookieFile)
	Handle(err, "Error al crear archivo de credenciales")
	defer f.Close()

	fmt.Println("Pega el valor de la cookie '.AspNetCore.Identity.Application':")
	var s string
	_, err = fmt.Scan(&s)
	Handle(err, "Error al leer entrada de usuario")
	_, err = f.Write([]byte(s))
	Handle(err, "Error al escribir contenido de archivo")
}
