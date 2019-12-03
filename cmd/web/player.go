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
	//pName := ""
	/*for _, cookie := range r.Cookies() {
		if strings.Contains(cookie.Value, "callsign") {
			pName = strings.ReplaceAll(cookie.Value, "callsign=", "")
		}
	}*/
	var cookie, err = r.Cookie("callsign")
	if err != nil {
		return ""
	}
	return cookie.Value
}
