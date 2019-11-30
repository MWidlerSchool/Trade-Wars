package main

import (
	"net/http"

	//"strconv"
	"html/template"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method == http.MethodPost {
		qStr = r.URL.Query()
		qStr := "Got a post in home: " + qStr.Get("xpos") + ", " + qStr.Get("ypos")
		http.Error(w, outStr, 200)
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
	}
}
