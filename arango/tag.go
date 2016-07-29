package arango

import (
	uuid "github.com/satori/go.uuid"
	"github.com/solher/arangolite"
	"github.com/tthanh/ims/model"
)

// TagStore implement model.TagStore interface
type TagStore struct {
	db *arangolite.DB
}

// NewTagStore create new TagStore
func NewTagStore(db *arangolite.DB) TagStore {
	return TagStore{
		db: db,
	}
}

// Create create new model.Tag
func (ts TagStore) Create(tag *model.Tag) error {
	tag.Key = uuid.NewV4().String()

	tx := arangolite.NewTransaction([]string{tagCollection}, []string{tagCollection}).
		AddQuery("newTag", `INSERT %v IN %v RETURN NEW._id`, toJSON(tag), tagCollection).Return("newTag")

	result := []string{}
	err := exec(ts.db, &result, tx)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return model.ErrNotExist
	}

	tag.ID = result[0]

	return nil
}

// GetAll get all model.Tag
func (ts TagStore) GetAll() ([]*model.Tag, error) {
	tx := arangolite.NewTransaction([]string{tagCollection}, nil).
		AddQuery("result", `FOR t IN %v RETURN t`, tagCollection).Return("result")

	result := []*model.Tag{}
	err := exec(ts.db, &result, tx)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, model.ErrNotExist
	}

	return result, nil
}

// GetByKey return model.Tag with given id
func (ts TagStore) GetByKey(key string) (*model.Tag, error) {
	tx := arangolite.NewTransaction([]string{tagCollection}, nil).
		AddQuery("result", `FOR t IN %v FILTER t._key==@key LIMIT 1 RETURN t`, tagCollection).
		Return("result").Bind("key", key)

	result := []*model.Tag{}
	err := exec(ts.db, &result, tx)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, model.ErrNotExist
	}

	return result[0], nil
}
