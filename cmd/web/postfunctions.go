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
	// alpha to int conversion returns two vals, so can't be done inline
	intX, _ := strconv.Atoi(x)
	intY, _ := strconv.Atoi(y)
	if shouldAllowMove(playerMap[playerName], intX, intY) {
		str1 := "updatePlayerLoc(" + x + ", " + y + ");"
		str2 := "setNavMessage('" + playerName + " sets course for sector " + x + ", " + y + "');"
		returnStr = str1 + str2
		// can't assign values to fields while in map
		p := playerMap[playerName]
		p.X = intX
		p.Y = intY
		playerMap[playerName] = p
	}
	return returnStr
}

// update player location on page without notification
func positionPlayer(playerName string) string {
	return "updatePlayerLoc(" + strconv.Itoa(playerMap[playerName].X) + ", " + strconv.Itoa(playerMap[playerName].Y) + ");"
}

func PostTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST received")
}


func otherPlayers() string{
	//otherPlayers := []string{"2","3"}
	return "updateOtherPlayersLoc(" + "2" + "," + "3" +");"

}