package main

import (
	//"log"
	"net/http"
)

type Player struct {
	X int
	Y int
}

func extractPlayerName(r *http.Request) string {
	var cookie, err = r.Cookie("Callsign")
	if err != nil {
		return ""
	}
	return cookie.Value
}
