package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	//"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
)

// global variable for player map
var playerMap = make(map[string]Player)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{}


type Message struct {
	Callsign string `json:"callsign"`
	Message  string `json:"message"`
}

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

	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	log.Println("http server started on port" + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
