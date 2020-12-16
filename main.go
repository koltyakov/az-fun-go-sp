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

	log.Printf("Go server has been started on http://127.0.0.1%s/", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func getAddr() string {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	return listenAddr
}
