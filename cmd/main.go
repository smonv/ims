package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/solher/arangolite"
	"github.com/tthanh/ims/arango"
	"github.com/tthanh/ims/message"
	"github.com/tthanh/ims/model"
	"github.com/tthanh/ims/server"

	"golang.org/x/net/websocket"
)

var (
	routes []func(req *message.Request) (*message.Response, error)
	s      *server.Server
)

func main() {
	db := arangolite.New().LoggerOptions(false, false, false).
		Connect("http://localhost:8529", "_system", "", "")

	arango.InitDatabase(db)

	imageStore := arango.NewImageStore(db)
	tagStore := arango.NewTagStore(db)
	imageTagStore := arango.NewImageTagStore(db)
	s = server.NewServer(imageStore, tagStore, imageTagStore)

	http.Handle("/ws", websocket.Handler(wsHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func wsHandler(ws *websocket.Conn) {
	var req message.Request
	for {
		websocket.JSON.Receive(ws, &req)

		resp, err := handle(s, &req)
		if err != nil {
			fmt.Println(err)
		}

		respMgs, err := json.Marshal(resp)
		if err != nil {
			fmt.Println(err)
		}

		websocket.JSON.Send(ws, respMgs)
	}
}

func handle(s *server.Server, req *message.Request) (*message.Response, error) {
	switch req.ActionType {
	case message.ActionTypeCreateTag:
		tag := &model.Tag{}
		err := json.Unmarshal(*req.ActionData, tag)
		if err != nil {
			return nil, err
		}
		return s.CreateTag(tag)
	}
	return nil, nil
}
