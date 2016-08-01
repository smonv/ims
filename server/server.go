package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tthanh/ims/model"
)

// Server struct
type Server struct {
	imageStore    model.ImageStore
	tagStore      model.TagStore
	imageTagStore model.ImageTagStore
}

// NewServer create new server
func NewServer(imageStore model.ImageStore, tagStore model.TagStore, imageTagStore model.ImageTagStore) *Server {
	return &Server{
		imageStore:    imageStore,
		tagStore:      tagStore,
		imageTagStore: imageTagStore,
	}
}

// Home ..
func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Hello")
}

func (s *Server) response(w http.ResponseWriter, v interface{}) {
	d, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(d)
}
