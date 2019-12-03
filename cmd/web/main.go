package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	//"github.com/gorilla/sessions"
)

// global variable for player map
var playerMap = make(map[string]Player)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/navigation", navigation)
	mux.HandleFunc("/trade", trade)
	mux.HandleFunc("/chat", chat)
	mux.HandleFunc("/players", playersHandler)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	godotenv.Load()
	var port = os.Getenv("PORT")

	log.Println("Starting server on :" + port)
	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
