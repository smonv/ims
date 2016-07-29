package arango

import (
	"testing"

	"github.com/tthanh/ims/model"
)

var (
	image = &model.Image{
		Name: "name",
		Size: 123456,
		Path: "path",
	}
)

func TestImageStore_Create(t *testing.T) {
	err := imageStore.Create(image)
	if err != nil {
		t.Fatal(err)
	}
	if len(image.ID) == 0 {
		t.Fatal("Image.ID empty")
	}
}

func TestImageStore_GetByKey(t *testing.T) {
	img, err := imageStore.GetByKey(image.Key)
	if err != nil {
		t.Fatal(err)
	}
	if img.Key != image.Key {
		t.Fatalf("Wrong image key: %s != %s", img.Key, image.Key)
	}
}

func TestImageStore_GetByName(t *testing.T) {
	img, err := imageStore.GetByName(image.Name)
	if err != nil {
		t.Fatal(err)
	}
	if img.Name != image.Name {
		t.Fatal("Get Wrong Image")
	}
}
