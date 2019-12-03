package main

import (
	"fmt"
	"net/http"
	//"os"
	//"github.com/gorilla/sessions"
)

func NavButtonPressed(x string, y string) string {
	str1 := "updatePlayerLoc(" + x + ", " + y + ");"
	str2 := "setNavMessage('Setting course for sector " + x + ", " + y + "');"
	return str1 + str2
	//outStr := string("Nav button pressed: " + x + ", " + y)
	//return outStr
}

func PostTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST received")
}
