package main

import (
	"net/http"
	"log"
	"simpleLogin/handlers"
	"simpleLogin/db"
)

func main() {
	handlers.InitHandlers()
	db.InitDB()
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

