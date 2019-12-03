package main

import (
	"net/http"
	"strings"
)

type Player struct {
	X int
	Y int
}

func extractPlayerName(r *http.Request) string {
	pName := ""
	for _, cookie := range r.Cookies() {
		if strings.Contains(cookie.Value, "callsign") {
			pName = strings.ReplaceAll(cookie.Value, "callsign=", "")
		}
	}
	return pName
}
