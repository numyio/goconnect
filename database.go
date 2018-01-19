package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// DB : Active Session
type DB struct {
	Session *mgo.Session
}

// Or other way to do it is use a single variable
// var db *mgo.Database

// Person : Persone Schema
type Person struct {
	Name  string
	Phone string
}

// Account : Account Schema
type Account struct {
	username string
	password string
}

// NewDB : MongoDB Session
func NewDB() *DB {
	fmt.Println("Connecting to database...")
	session, err := mgo.Dial("localhost") // connect to mongo
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
	session.SetMode(mgo.Monotonic, true)
	return &DB{Session: session}
}
