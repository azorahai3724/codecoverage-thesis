package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	println("Serving")

	router.Handle("/", http.FileServer(http.Dir("/Users/kananeyvazov/codecoverage-thesis/frontend")))
	router.HandleFunc("/app/create", newApp).Methods("POST", "OPTIONS")
	//router.HandleFunc("/app/update", newApp).Methods("POST")
	router.HandleFunc("/app", getAllApps).Methods("GET")
	router.HandleFunc("/app/{Name}", getOneApp).Methods("GET")

	if err := http.ListenAndServe(":10000", router); err != nil {
		log.Fatal("Listen and serve: %w", err)
	}

}
