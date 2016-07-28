package arango

import (
	"github.com/solher/arangolite"
	"github.com/tthanh/ims/model"
)

// ImageStore implement model.ImageStore interface
type ImageStore struct {
	db *arangolite.DB
}

// NewImageStore return new ImageStore
func NewImageStore(db *arangolite.DB) ImageStore {
	return ImageStore{
		db: db,
	}
}

// CreateImage create new image
func (is ImageStore) CreateImage(image *model.Image) error {
	tx := arangolite.NewTransaction([]string{imageCollection}, []string{imageCollection}).
		AddQuery("newImage", `INSERT %v IN %v`, toJSON(image), imageCollection)

	_, err := is.db.Run(tx)
	if err != nil {
		return err
	}
}
