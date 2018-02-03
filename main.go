package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// User : for input form
type User struct {
	Name     string `json:"name" form:"name" query:"name"`
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

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

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// newUser : Uses User struct for new users
func newUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	db := newDB().Session
	cf := db.DB("goconnect").C("accounts")
	cf.Insert(u)
	result := c.JSON(http.StatusCreated, u)
	fmt.Println("Inserting Username: " + u.Name)
	fmt.Println("Inserting Password: " + u.Password)
	fmt.Println("Inserting Email: " + u.Email)
	return result
}
