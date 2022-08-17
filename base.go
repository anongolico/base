package base

import (
	"net/http"
)

const (
	BaseUrl      = "https://rouzed.one/Hilo/"
	BaseMediaUrl = "https://rouzed.one/Media/"
)

// Client used to make the http requests
var Client = http.Client{}
