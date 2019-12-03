package main

import (
	"fmt"
	"net/http"
	//"os"
	//"github.com/gorilla/sessions"
)

func NavButtonPressed(playername string, x string, y string) string {
	str1 := "updatePlayerLoc(" + x + ", " + y + ");"
	str2 := "setNavMessage('" + playername + " sets course for sector " + x + ", " + y + "');"
	return str1 + str2
}

func PostTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST received")
}
