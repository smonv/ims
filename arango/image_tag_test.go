package arango

import (
	"fmt"
	"testing"

	"github.com/tthanh/ims/model"
)

var (
	ri = &model.Image{
		Name: "rImage",
		Size: 123456,
		Path: "path",
	}
	rt = &model.Tag{
		Name: "rTag",
		Path: "path1",
	}
)

func TestImageTagStore_Create(t *testing.T) {
	err := imageStore.Create(ri)
	if err != nil {
		fmt.Println(err)
	}

	err = tagStore.Create(rt)
	if err != nil {
		fmt.Println(err)
	}

	it := &model.ImageTag{
		From: ri.ID,
		To:   rt.ID,
	}

	err = imageTagStore.Create(it)
	if err != nil {
		t.Fatal(err)
	}
}

func TestImageTagStore_GetImages(t *testing.T) {
	images, err := imageTagStore.GetImages(rt.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(images) == 0 {
		t.Fatal(model.ErrNotExist)
	}
}

func TestImageTagStore_GetTags(t *testing.T) {
	tags, err := imageTagStore.GetTags(ri.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(tags) == 0 {
		t.Fatal(model.ErrNotExist)
	}
}
