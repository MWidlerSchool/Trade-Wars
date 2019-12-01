package main

import (
	"fmt"
	"log"
	"net/http"
	//"os"
	"github.com/gorilla/sessions"
)

func NavButtonPressed(xLoc string, yLoc string) {
	log.Println("x = " + xLoc + ", y = " + yLoc)
}

func PostTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST received")
}
