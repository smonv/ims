package server

import "github.com/tthanh/ims/model"

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
