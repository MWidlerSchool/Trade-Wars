package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/navigation", navigation)
	mux.HandleFunc("/trade", trade)
	mux.HandleFunc("/chat", chat)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	var port = os.Getenv("PORT")

	log.Println("Starting server on :" + port)
	err := http.ListenAndServe(":" + port, mux)
	log.Fatal(err)
}
