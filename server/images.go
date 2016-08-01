package server

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// GetImage ...
func (s *Server) GetImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v := vars["key"]
	if len(v) == 0 {
		err := errors.New("Key not found")
		s.response(w, err)
	}

	image, err := s.imageStore.GetByKey(v)
	if err != nil {
		image, err = s.imageStore.GetByName(v)
		if err != nil {
			s.response(w, err)
		}
	}

	s.response(w, image)
}
