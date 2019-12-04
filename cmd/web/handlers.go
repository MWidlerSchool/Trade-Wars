package main

import (
	"net/http"

	//"strconv"
	"html/template"
	"log"

	//"github.com/gorilla/sessions"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method == http.MethodPost {
		playerName := extractPlayerName(r)
		qStr := r.URL.Query()
		switch qStr.Get("actiontype") {
		case "navbutton":
			w.Header().Set("Content-Type", "application/javascript")
			w.Write([]byte(NavButtonPressed(playerName, qStr.Get("xpos"), qStr.Get("ypos"))))
		case "navloaded":
			w.Header().Set("Content-Type", "application/javascript")
			w.Write([]byte(positionPlayer(playerName)))
		}
		return
	}
	ts, err := template.ParseFiles("./ui/html/tradewars.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func navigation(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/navigation" {
		http.NotFound(w, r)
		return
	}
	if r.Method == http.MethodPost {
		outStr := "Got a post in nav: " + r.FormValue("xPos") + ", " + r.FormValue("yPos")
		http.Error(w, outStr, 200)
		return
	}
	ts, err := template.ParseFiles("./ui/html/navigation.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func trade(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/trade" {
		http.NotFound(w, r)
		return
	}
	ts, err := template.ParseFiles("./ui/html/trade.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func playersHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/players" {
		http.NotFound(w, r)
		return
	}
	ts, err := template.ParseFiles("./ui/html/players.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
		callsign := r.Form.Get("callsign")
		cookie := http.Cookie{
			Name:    "callsign",
			Value:   callsign,
			Expires: time.Now().AddDate(0, 0, 1),
			Path:    "/",
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/map", http.StatusSeeOther)
	}
}

func mapHandler(w http.ResponseWriter, r *http.Request) {
	var cookie, err = r.Cookie("callsign")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error, Could not Process callsign from cookie", 50)
		return
	}
	callsign := cookie.Value
	log.Println(callsign)
	w.Write([]byte(callsign))
	// add character if not already present
	_, isPresent := playerMap[callsign]
	if isPresent == false {
		playerMap[callsign] = Player{X: 4, Y: 4}
	}
}

func chat(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/chat" {
		http.NotFound(w, r)
		return
	}
	ts, err := template.ParseFiles("./ui/html/chat.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	clients[ws] = true
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}