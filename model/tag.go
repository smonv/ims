package model

// Tag struct
type Tag struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
	Path string `json:"path"`
	UUID string `json:"uuid"`
}

// TagStore interface
type TagStore interface {
	CreateTag(tag *Tag) error

	GetTags() ([]*Tag, error)
	GetTagByUUID(uuid string) (*Tag, error)
}
