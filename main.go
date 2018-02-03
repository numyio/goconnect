package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	go echoServer()
	httpServer()

}

func echoServer() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "/static")
	fmt.Println("echo")

	log.Fatal(e.Start(":1323"))

}

func httpServer() {
	fmt.Println("Listening on port 8000 for REST requests")
	router := mux.NewRouter()
	router.HandleFunc("/accounts", getAccounts).Methods("GET")
	router.HandleFunc("/accounts/{id}", getAccount).Methods("GET")
	router.HandleFunc("/accounts/{id}", createAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}", deleteAccount).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
