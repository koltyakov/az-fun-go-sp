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
	http.HandleFunc("/api/GetLists", h.GetLists)
	http.HandleFunc("/api/GetFields", h.GetFields)
	http.HandleFunc("/api/Storage", h.Storage)
}
