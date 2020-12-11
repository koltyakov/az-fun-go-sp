package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	listenAddr := getAddr()

	http.HandleFunc("/api/GetLists", getListsHandler)
	http.HandleFunc("/api/GetFields", getFieldsHandler)

	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func getAddr() string {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	return listenAddr
}
