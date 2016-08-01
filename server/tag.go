package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tthanh/ims/model"
)

// CreateTag ...
func (s *Server) CreateTag(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	tag := &model.Tag{}
	err := decoder.Decode(tag)
	if err != nil {
		s.response(w, err)
	}

	err = s.tagStore.Create(tag)
	if err != nil {
		s.response(w, err)
	}

	s.response(w, tag)
}

// GetTags ...
func (s *Server) GetTags(w http.ResponseWriter, r *http.Request) {
	tags, err := s.tagStore.GetAll()
	if err != nil {
		s.response(w, err)
	}
	s.response(w, tags)
}

// GetTag ...
func (s *Server) GetTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	tag, err := s.tagStore.GetByKey(key)
	if err != nil {
		s.response(w, err)
	}
	s.response(w, tag)
}
