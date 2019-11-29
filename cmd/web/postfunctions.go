package main

import (
	"log"
	"net/http"
	"os"
)

func navButtonPressed(xLoc int, yLoc int) {
	log.Println("x = " + xLoc + ", y = " + yLoc)
}