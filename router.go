package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func getAccounts(w http.ResponseWriter, r *http.Request) {

	// result := Person{}

	fmt.Println("Accounts request")

}

func getAccount(w http.ResponseWriter, r *http.Request) {
	db := NewDB()
	c := db.Session.DB("goconnect").C("people")
	result := Person{}
	c.Find(bson.M{"name": "Ale"}).One(&result)
	c.Find(bson.M{"name": "Bob"}).One(&result)

	fmt.Println("Phone:", result.Phone)
	fmt.Println("Account request")
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	// c.Insert(
	// 	&Person{"Ale", "3321244323"},
	// 	&Person{"Bob", "8081122222"},
	// )
	fmt.Println("Account creation")
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Account Deletion")
}
