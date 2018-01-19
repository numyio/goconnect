package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Listening on port 8000 for REST requests")
	router := mux.NewRouter()
	router.HandleFunc("/accounts", getAccounts).Methods("GET")
	router.HandleFunc("/accounts/{id}", getAccount).Methods("GET")
	router.HandleFunc("/accounts/{id}", createAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}", deleteAccount).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}
