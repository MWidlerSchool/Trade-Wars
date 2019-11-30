package main

import (
	"log"
	//"net/http"
	//"os"
)

func NavButtonPressed(xLoc string, yLoc string) {
	log.Println("x = " + xLoc + ", y = " + yLoc)
}

func PostTest(w *http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST received")
}
