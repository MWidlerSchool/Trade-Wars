package main

import (
	"fmt"
	"net/http"
	//"os"
)

func NavButtonPressed(x string, y string) string {
	outStr := string("Got a post in home: " + x + ", " + y)
	return outStr
}

func PostTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST received")
}
