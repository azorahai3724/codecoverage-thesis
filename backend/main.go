package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	println("Serving")

	router.HandleFunc("/app/create", newReport).Methods("POST", "OPTIONS")
	router.HandleFunc("/app", getAllApps).Methods("GET")
	router.HandleFunc("/app/{Name}", getOneApp).Methods("GET")

	err := http.ListenAndServe(":10000", router)
	if err != nil {
		log.Fatal("Listen and serve: %w", err)
	}
}
