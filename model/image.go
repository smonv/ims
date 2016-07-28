package model

// Image struct
type Image struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
	Size uint64 `json:"size"`
	Path string `json:"path"`
}

// ImageStore interface
type ImageStore interface {
	CreateImage(image *Image) error
}
