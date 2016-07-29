package main

import (
	"fmt"
	"net/http"

	"github.com/tthanh/ims/message"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/ws", websocket.Handler(handler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handler(ws *websocket.Conn) {
	var request message.Request
	for {
		websocket.JSON.Receive(ws, &request)
		fmt.Printf("%v\n", request)
	}
}
