package main

import (
	"fmt"
	"net/http"
	"strconv"
	//"os"
	//"github.com/gorilla/sessions"
)

func NavButtonPressed(playerName string, x string, y string) string {
	returnStr := "setNavMessage('Target sector must be adjacent to current location!');"
	intX, _ := strconv.Atoi(x)
	intY, _ := strconv.Atoi(y)
	if shouldAllowMove(playerMap[playerName], intX, intY) {
		str1 := "updatePlayerLoc(" + x + ", " + y + ");"
		str2 := "setNavMessage('" + playerName + " sets course for sector " + x + ", " + y + "');"
		returnStr = str1 + str2
		p := playerMap[playerName]
		p.X = intX
		p.Y = intY
		playerMap[playerName] = p
	}
	return returnStr
}

func PostTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST received")
}
