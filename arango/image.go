package arango

import (
	uuid "github.com/satori/go.uuid"
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

// Create create new image
func (is ImageStore) Create(image *model.Image) error {
	image.Key = uuid.NewV4().String()

	tx := arangolite.NewTransaction([]string{imageCollection}, []string{imageCollection}).
		AddQuery("newImage", `INSERT %v IN %v RETURN NEW._id`, toJSON(image), imageCollection).
		Return("newImage")

	result := []string{}
	err := exec(is.db, &result, tx)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return model.ErrNotExist
	}

	image.ID = result[0]

	return nil
}

// GetByKey return model.Image with given key
func (is ImageStore) GetByKey(key string) (*model.Image, error) {
	tx := arangolite.NewTransaction([]string{imageCollection}, nil).
		AddQuery("result", `FOR i IN %v FILTER i._key==@key LIMIT 1 RETURN i`, imageCollection).
		Return("result").Bind("key", key)

	result := []*model.Image{}
	err := exec(is.db, &result, tx)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, model.ErrNotExist
	}

	return result[0], nil
}

// GetByName return model.Image with given name
func (is ImageStore) GetByName(name string) (*model.Image, error) {
	tx := arangolite.NewTransaction([]string{imageCollection}, nil).
		AddQuery("result", `FOR i IN %v FILTER i.name==@name LIMIT 1 RETURN i`, imageCollection).
		Return("result").Bind("name", name)

	result := []*model.Image{}
	err := exec(is.db, &result, tx)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, model.ErrNotExist
	}

	return result[0], nil
}
