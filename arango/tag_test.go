package arango

import (
	"testing"

	"github.com/tthanh/ims/model"
)

var (
	t1 = &model.Tag{
		Name: "tag1",
		Path: "path1",
		UUID: "1",
	}

	t2 = &model.Tag{
		Name: "tag2",
		Path: "path2",
		UUID: "2",
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
	if len(result) == 0 {
		t.Fatalf("len: %d = 0", len(result))
	}
}

func TestGetTagByUUID(t *testing.T) {
	tag, err := tagStore.GetTagByUUID(t1.UUID)
	if err != nil {
		t.Fatal(err)
	}
	if tag.UUID != t1.UUID {
		t.Fatalf("%v != %v\n", tag, t1)
	}
}
