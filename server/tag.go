package server

import (
	"github.com/tthanh/ims/message"
	"github.com/tthanh/ims/model"
)

// CreateTag ...
func (s *Server) CreateTag(tag *model.Tag) (*message.Response, error) {
	err := s.tagStore.Create(tag)
	if err != nil {
		return nil, err
	}

	return &message.Response{
		Data: tag,
	}, nil
}
