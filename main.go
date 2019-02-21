package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

type wraper struct {
}

func main() {
	interactKraken()
}

func interactKraken() {
	connection, err := websocket.Dial("wss://ws-sandbox.kraken.com/ping", "wss", "wss://ws-sandbox.kraken.com")
	if err != nil {
		log.Printf("%+v", err)
	}
	message := `{
		"event": "ping",
		"reqid": 42
	  }`
	fmt.Println(connection, message)
	// numBytes, err := connection.Write([]byte(message))
	// if err != nil {
	// 	log.Printf("%+v", err)
	// } else {
	// 	log.Printf("%+v", numBytes)
	// }
	// if err != nil {
	// 	log.Printf("%+v", err)
	// }
	// fmt.Println(res)
}
