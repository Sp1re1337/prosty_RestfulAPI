package handlers

import (
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
)

func WebsocketHandler(c *websocket.Conn) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in WebSocketHandler", r)
		}
		c.Close()
	}()

	for {
		messageType, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		fmt.Printf("received: %s\n", msg)

		if err = c.WriteMessage(messageType, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
}