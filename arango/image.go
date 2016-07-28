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
	return nil
}

// GetImageByName return model.Image with given name
func (is ImageStore) GetImageByName(name string) (*model.Image, error) {
	tx := arangolite.NewTransaction([]string{imageCollection}, nil).
		AddQuery("var1", `FOR i IN %v FILTER i.name==@key LIMIT 1 RETURN i`, imageCollection).
		Return("var1").Bind("key", name)

	var result []*model.Image
	err := exec(is.db, &result, tx)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, model.ErrNotExist
	}

	return result[0], nil
}
