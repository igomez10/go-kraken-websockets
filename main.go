package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type wraper struct {
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	interactKraken()
}

func interactKraken() {

	log.Print("\n\n\n\n\n\n\n\n\n\n\n\n")
	connection, err := createConnectionToKraken()
	if err != nil {
		log.Println("Could not connect")
	}

	payload := `{
		"event": "subscribe",
		"pair": [
		  "XBT/USD",
		  "XBT/EUR"
		],
		"subscription": {
		  "name": "*"
		}
	  }`

	writeToSocket(connection, []byte(payload))

	address := connection.RemoteAddr()
	log.Println(address.String())
	arr := make([]int, 7)

	for range arr {
		incomingBytes := readFromSocket(connection)
		log.Printf("Received message: %s", incomingBytes)
		time.Sleep(time.Millisecond * 200)
	}
}

func createConnectionToKraken() (*websocket.Conn, error) {
	websocketHost := "wss://ws.kraken.com"
	// protocol := "wss"

	connection, response, err := websocket.DefaultDialer.Dial(websocketHost, nil)
	if err != nil {
		log.Printf("ERROR creating connection to url, %v", err)
	} else {
		log.Printf("SUCCESS: Connection established with %s  \n", websocketHost)

		if err != nil {
			log.Println("Error marshaling json:", err)
		}
		log.Printf("RESPONSE: %+v", *response)
	}
	return connection, err
}

func writeToSocket(connection *websocket.Conn, payload []byte) error {

	err := connection.WriteJSON([]byte(payload))
	if err != nil {
		log.Printf("%+v \n", err)
	} else {
		log.Printf("SUCCESS: Wrote to the sockets")
	}
	return err
}

func readFromSocket(connection *websocket.Conn) interface{} {
	//		arr := make([]byte, 500)
	var res interface{}
	err := connection.ReadJSON(res)
	if err != nil {
		log.Printf("ERROR: Could not read from connection, %+v \n", err)
	} else {
		log.Printf("Success reading bytes from connection \n")
	}
	return res
}
