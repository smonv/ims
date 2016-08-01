package server

import (
	"encoding/json"
	"html/template"
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

// Index ...
func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("www/index.html")
	if err != nil {
		s.response(w, err)
	}
	t.Execute(w, nil)
}

func (s *Server) response(w http.ResponseWriter, v interface{}) {
	d, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(d)
}
