package server

import (
	"encoding/json"
	"net/http"

	"github.com/tthanh/ims/model"
)

// CreateTag ...
func (s *Server) CreateTag(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	tag := &model.Tag{}
	err := decoder.Decode(tag)
	if err != nil {
		panic(err)
	}

	err = s.tagStore.Create(tag)
	if err != nil {
		panic(err)
	}

	response(w, tag)
}

// GetTags ...
func (s *Server) GetTags(w http.ResponseWriter, r *http.Request) {
	tags, err := s.tagStore.GetAll()
	if err != nil {
		response(w, err)
	}
	response(w, tags)
}
