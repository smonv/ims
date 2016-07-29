package model

// Image struct
type Image struct {
	Key  string `json:"_key"`
	ID   string `json:"_id"`
	Name string `json:"name"`
	Size uint64 `json:"size"`
	Path string `json:"path"`
}

// ImageStore interface
type ImageStore interface {
	Create(image *Image) error

	GetByKey(key string) (*Image, error)
	GetByName(name string) (*Image, error)
}
