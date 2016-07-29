package arango

import (
	"encoding/json"

	"github.com/solher/arangolite"
)

const (
	graphName          = "ims"
	imageCollection    = "image"
	tagCollection      = "tag"
	imageTagCollection = "image_tag"
)

func toJSON(v interface{}) string {
	r, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(r)
}

func exec(db *arangolite.DB, result interface{}, query arangolite.Runnable) error {
	data, err := db.Run(query)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		panic(err)
	}
	return nil
}
