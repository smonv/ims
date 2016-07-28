package arango

import (
	"reflect"
	"testing"

	"github.com/tthanh/ims/model"
)

var (
	t1 = &model.Tag{
		ID:   "1",
		Name: "tag1",
		Path: "path1",
	}

	t2 = &model.Tag{
		ID:   "2",
		Name: "tag2",
		Path: "path2",
	}
)

func TestCreateTag(t *testing.T) {
	err := tagStore.CreateTag(t1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTags(t *testing.T) {
	_ = tagStore.CreateTag(t2)

	result, err := tagStore.GetTags()
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 2 {
		t.Fatalf("len: %d != %d", len(result), 2)
	}
}

func TestGetTagByID(t *testing.T) {
	tag, err := tagStore.GetTagByID(t1.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(tag, t1) {
		t.Fatalf("%v != %v\n", tag, t1)
	}
}
