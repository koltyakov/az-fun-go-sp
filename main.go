package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	initRoutes()
	listenAddr := getAddr()
	log.Printf("Custom handlers server is running on http://127.0.0.1%s", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func getAddr() string {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	return listenAddr
}
