package main

import (
	"net/http"
)

type Player struct {
	X int
	Y int
}

// get the player name from the cookie
func extractPlayerName(r *http.Request) string {
	var cookie, err = r.Cookie("Callsign")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// check if a move is legal
func shouldAllowMove(p Player, x int, y int) bool {
	if p.X-x <= 1 && p.X-x >= -1 && p.Y-y <= 1 && p.Y-y >= -1 {
		return true
	}
	return false
}
