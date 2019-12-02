package main

import (
	"fmt"
	"net/http"
	//"os"
	//"github.com/gorilla/sessions"
)

func NavButtonPressed(x string, y string) string {
	return "updatePlayerLoc(" + x + ", " + y + ");"
	//outStr := string("Nav button pressed: " + x + ", " + y)
	//return outStr
}

func PostTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST received")
}
