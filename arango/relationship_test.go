package arango

import (
	"fmt"
	"testing"

	"github.com/tthanh/ims/model"
)

var (
	imageName = "imageRelationship"
	tagUUID   = "tagRelationship"
	it        *model.ImageTag
	rTag      *model.Tag
	rImage    *model.Image
)

func init() {
	i := &model.Image{
		Name: imageName,
		Size: 123456,
		Path: "path",
	}

	t := &model.Tag{
		Name: "tag1",
		Path: "path1",
		UUID: tagUUID,
	}

	err := imageStore.Create(i)
	if err != nil {
		fmt.Println(err)
	}

	err = tagStore.CreateTag(t)
	if err != nil {
		fmt.Println(err)
	}

	rTag, err = tagStore.GetTagByUUID(tagUUID)
	if err != nil {
		fmt.Println(err)
	}
	rImage, err = imageStore.GetImageByName(imageName)
	if err != nil {
		fmt.Println(err)
	}
}

func TestCreate(t *testing.T) {
	it = &model.ImageTag{
		From: rImage.ID,
		To:   rTag.ID,
	}

	err := imageTagStore.Create(it)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetImages(t *testing.T) {
	images, err := imageTagStore.GetImages(rTag.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(images) == 0 {
		t.Fatal(model.ErrNotExist)
	}
}

func TestRelationshipGetTags(t *testing.T) {
	tags, err := imageTagStore.GetTags(rImage.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(tags) == 0 {
		t.Fatal(model.ErrNotExist)
	}
}
