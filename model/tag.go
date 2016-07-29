package model

// Tag struct
type Tag struct {
	Key  string `json:"_key"`
	ID   string `json:"_id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

// TagStore interface
type TagStore interface {
	Create(tag *Tag) error

	GetAll() ([]*Tag, error)
	GetByKey(key string) (*Tag, error)
}
