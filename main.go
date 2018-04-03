package main

import (
	"net/http"
	"log"
	"github.com/dnlo/web/simpleLogin/handlers"
	"github.com/dnlo/web/simpleLogin/db"
)

func main() {
	handlers.InitHandlers()
	db.InitDB()
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

