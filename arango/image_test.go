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

func TestCreateImage(t *testing.T) {
	err := imageStore.CreateImage(image)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetImageByName(t *testing.T) {
	img, err := imageStore.GetImageByName(image.Name)
	if err != nil {
		t.Fatal(err)
	}
	if img.Name != image.Name {
		t.Fatal("Get Wrong Image")
	}
}
