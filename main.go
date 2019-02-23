package main

import (
	"log"
	"time"

	"golang.org/x/net/websocket"
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
	for _ = range make([]int, 7) {

		incomingBytes := readFromSocket(connection)
		log.Printf("Received message: %s", incomingBytes)
		time.Sleep(time.Millisecond * 200)
	}

}

func createConnectionToKraken() (*websocket.Conn, error) {
	websocketHost := "wss://ws.kraken.com"
	// protocol := "wss"
	connection, err := websocket.Dial(websocketHost, "", "")
	if err != nil {
		log.Printf("ERROR creating connection to url, %v", err)
	} else {
		log.Printf("SUCCESS: Connection established with %s  \n", websocketHost)
	}
	return connection, err
}

func writeToSocket(connection *websocket.Conn, payload []byte) int {
	numWrites, err := connection.Write([]byte(payload))
	if err != nil {
		log.Printf("%+v \n", err)
	} else {
		//log.Printf("SUCCESS: Wrote %d bytes to the sockets '%s'", numWrites, payload)
	}
	return numWrites
}

func readFromSocket(connection *websocket.Conn) []byte {
	arr := make([]byte, 500)
	numReads, err := connection.Read(arr)
	if err != nil {
		log.Printf("ERROR: Could not read from connection, %+v \n", err)
	} else {
		log.Printf("Success reading %d bytes from connection \n", numReads)
	}
	return arr
}
