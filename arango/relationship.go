package arango

import (
	"github.com/solher/arangolite"
	"github.com/tthanh/ims/model"
)

// ImageTagStore implement model.ImageTag interface
type ImageTagStore struct {
	db *arangolite.DB
}

// NewImageTagStore create new ImageTagStore
func NewImageTagStore(db *arangolite.DB) ImageTagStore {
	return ImageTagStore{
		db: db,
	}
}

// Create create new model.ImageTag
func (its ImageTagStore) Create(it *model.ImageTag) error {
	tx := arangolite.NewTransaction([]string{imageTagCollection}, []string{imageTagCollection}).
		AddQuery("newImageTag", `INSERT %v IN %v`, toJSON(it), imageTagCollection)

	_, err := its.db.Run(tx)
	if err != nil {
		return err
	}

	return nil
}

// GetImages return Images with given tag
func (its ImageTagStore) GetImages(tagID string) ([]*model.Image, error) {
	tx := arangolite.NewTransaction([]string{imageTagCollection, imageCollection}, nil).
		AddQuery("var1", `FOR i IN INBOUND @key GRAPH @graph RETURN i`).
		Return("var1").Bind("key", tagID).Bind("graph", graphName)

	var result []*model.Image
	err := exec(its.db, &result, tx)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, model.ErrNotExist
	}
	return result, nil
}

// GetTags return Tags with given Image
func (its ImageTagStore) GetTags(imageID string) ([]*model.Tag, error) {
	tx := arangolite.NewTransaction([]string{imageTagCollection, tagCollection}, nil).
		AddQuery("var1", `FOR t IN OUTBOUND @key GRAPH @graph RETURN t`).
		Return("var1").Bind("key", imageID).Bind("graph", graphName)

	var result []*model.Tag
	err := exec(its.db, &result, tx)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, model.ErrNotExist
	}

	return result, nil
}
