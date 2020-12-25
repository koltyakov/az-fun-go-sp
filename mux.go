package main

import (
	"log"
	"net/http"

	"github.com/koltyakov/az-fun-go-sp/handlers"
)

// Routes configuration
func initRoutes() {
	// Creating SharePoint API client instance
	sp := getSP()

	// Creating Azure Storage Account API instance
	sa, err := getSA()
	if err != nil {
		log.Fatalf("can't create storage account client, %s\n", err)
	}

	// Constructing handlers struct
	h := handlers.NewHandlers(sp, sa)

	// Binding Functions with Handlers
	http.HandleFunc("/api/lists", h.Lists)
	http.HandleFunc("/api/fields", h.Fields)
	http.HandleFunc("/api/storage", h.Storage)

	// Wildcard route /api/*
	http.HandleFunc("/api/", h.Wildcard)

	// Timer job(s)
	http.HandleFunc("/timer", h.Timer)
}
