package main

import (
	"net/http"

	"github.com/koltyakov/az-fun-go-sp/handlers"
)

// Routes configuration
func initRoutes() {
	h := handlers.NewHandlers(getSP())

	http.HandleFunc("/api/GetLists", h.GetLists)
	http.HandleFunc("/api/GetFields", h.GetFields)
}
