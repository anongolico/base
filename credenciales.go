package base

import (
	"net/http"
)

var (
	AccessCookie = http.Cookie{
		Name:     ".AspNetCore.Identity.Application",
		Value:    "",
		Secure:   true,
		HttpOnly: true,
	}
)

func SetAccessCookieValue(v string) {
	AccessCookie.Value = v
}
