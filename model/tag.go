package model

// Tag struct
type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

// TagStore interface
type TagStore interface {
	CreateTag(tag *Tag) error

	GetTags() ([]*Tag, error)
	GetTagByID(id string) (*Tag, error)
}
