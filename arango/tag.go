package arango

import (
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

// CreateTag create new model.Tag
func (ts TagStore) CreateTag(tag *model.Tag) error {
	tx := arangolite.NewTransaction([]string{tagCollection}, []string{tagCollection}).
		AddQuery("newTag", `INSERT %v IN %v`, toJSON(tag), tagCollection)

	_, err := ts.db.Run(tx)
	if err != nil {
		return err
	}

	return nil
}

// GetTags get all model.Tag
func (ts TagStore) GetTags() ([]*model.Tag, error) {
	tx := arangolite.NewTransaction([]string{tagCollection}, nil).
		AddQuery("var1", `FOR t IN %v RETURN t`, tagCollection).Return("var1")

	var result []*model.Tag
	err := exec(ts.db, &result, tx)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, model.ErrNotExist
	}

	return result, nil
}

// GetTagByUUID return model.Tag with given id
func (ts TagStore) GetTagByUUID(uuid string) (*model.Tag, error) {
	tx := arangolite.NewTransaction([]string{tagCollection}, nil).
		AddQuery("var1", `FOR t IN %v FILTER t.uuid==@key LIMIT 1 RETURN t`, tagCollection).
		Return("var1").Bind("key", uuid)

	var result []*model.Tag
	err := exec(ts.db, &result, tx)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, model.ErrNotExist
	}

	return result[0], nil
}
