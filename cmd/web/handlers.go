package main

import (
	"net/http"

	//"strconv"
	"html/template"
	"log"
	"github.com/gorilla/sessions"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method == http.MethodPost {
		qStr := r.URL.Query()
		returnStr := ""
		// main switch for incoming POST requests
		switch qStr.Get("actiontype") {
		case "navbutton":
			returnStr = NavButtonPressed(qStr.Get("xpos"), qStr.Get("ypos"))
		}
		http.Error(w, returnStr, 200)
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
	}else if r.Method == http.MethodPost{
		err := r.ParseForm()
		if err != nil{
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
		callsign :=r.Form.Get("callsign")
		cookie := http.Cookie{
			Name: "callsign",
			Value: callsign,
			Expires: time.Now.AddDate(0,0,1),
			Path:"/"
		}
		http.Redirect(w, r, "/map", http.StatusSeeOther)
	}
}

func mapHandler(w http.ResponseWriter, r *http.Request){
	var cookie, err = r.Cookie("callsign")
	if err != nil{
		log.Println(err.Error())
		http.Error(w, "Internal Server Error, Could not Process callsign from cookie", 50)
		return
	}
	callsign := cookie.Valuelog.Println(callsign)
	w.Write([]byte(callsign))
}