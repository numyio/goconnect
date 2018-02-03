package main

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	upgrader = websocket.Upgrader{}
)

func sendAndRecieve(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	defer ws.Close()

	for {
		// send to client
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Receive from server
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)

	}
}
