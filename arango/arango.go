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

// InitDatabase ...
func InitDatabase(db *arangolite.DB) {
	_, _ = db.Run(&arangolite.CreateDatabase{Name: "ims"})

	db.SwitchDatabase("ims")

	for _, collection := range []string{imageCollection, tagCollection} {
		_, _ = db.Run(&arangolite.CreateCollection{Name: collection})
	}

	_, err := db.Run(&arangolite.GetGraph{Name: "ims_graph"})
	if err != nil {
		from := make([]string, 1)
		from[0] = imageCollection
		to := make([]string, 1)
		to[0] = tagCollection
		edgeDefinition := arangolite.EdgeDefinition{Collection: imageTagCollection, From: from, To: to}
		edgeDefinitions := make([]arangolite.EdgeDefinition, 1)
		edgeDefinitions[0] = edgeDefinition
		db.Run(&arangolite.CreateGraph{Name: "ims_graph", EdgeDefinitions: edgeDefinitions})
	}
}

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
