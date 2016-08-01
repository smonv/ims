package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tthanh/ims/model"
)

// CreateRelationship ...
func (s *Server) CreateRelationship(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	it := &model.ImageTag{}
	err := decoder.Decode(it)
	if err != nil {
		panic(err)
	}

	err = s.imageTagStore.Create(it)
	if err != nil {
		s.response(w, err)
	}

	s.response(w, it)
}

// GetImagesByTag ...
func (s *Server) GetImagesByTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	if len(key) == 0 {
		s.response(w, errors.New("Key not found"))
	}

	images, err := s.imageTagStore.GetImages(key)
	if err != nil {
		s.response(w, err)
	}

	s.response(w, images)
}

// GetTagsByImage ...
func (s *Server) GetTagsByImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	if len(key) == 0 {
		s.response(w, errors.New("Key not found"))
	}

	tags, err := s.imageTagStore.GetTags(key)
	if err != nil {
		s.response(w, err)
	}

	s.response(w, tags)
}
