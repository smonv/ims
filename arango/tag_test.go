package arango

import (
	"testing"

	"github.com/tthanh/ims/model"
)

var (
	t1 = &model.Tag{
		Name: "tag1",
		Path: "path1",
	}

	t2 = &model.Tag{
		Name: "tag2",
		Path: "path2",
	}
)

func TestTagStore_Create(t *testing.T) {
	err := tagStore.Create(t1)
	if err != nil {
		t.Fatal(err)
	}

	if len(t1.ID) == 0 {
		t.Fatal("New Tag.ID empty")
	}
}

func TestTagStore_GetAll(t *testing.T) {
	_ = tagStore.Create(t2)

	result, err := tagStore.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(result) == 0 {
		t.Fatalf("len: %d = 0", len(result))
	}
}

func TestTagStore_GetByKey(t *testing.T) {
	tag, err := tagStore.GetByKey(t1.Key)
	if err != nil {
		t.Fatal(err)
	}
	if tag.Key != t1.Key {
		t.Fatalf("%v != %v\n", tag, t1)
	}
}
