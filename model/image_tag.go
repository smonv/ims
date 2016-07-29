package model

// ImageTag relationship
type ImageTag struct {
	From string `json:"_from"`
	To   string `json:"_to"`
}

// ImageTagStore interface
type ImageTagStore interface {
	Create(it *ImageTag) error

	GetImages(tagID string) ([]*Image, error)
	GetTags(imageID string) ([]*Tag, error)
}
