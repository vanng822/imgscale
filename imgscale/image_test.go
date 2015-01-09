package imgscale

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"path/filepath"
)

func TestGetImageWrongFile(t *testing.T) {
	filename, _ := filepath.Abs("./data/kth.jpg")
	info := &ImageInfo{filename, 100, 0.0, "jpg", false, ""}
	_, err := GetImage(info)
	assert.Error(t, err)
}

func TestGetImageScaleOK(t *testing.T) {
	filename, _ := filepath.Abs("../data/kth.jpg")
	info := &ImageInfo{filename, 100, 0.0, "jpg", false, ""}
	img, err := GetImage(info)
	assert.Equal(t, 100, img.GetImageHeight())
	assert.Equal(t, 133, img.GetImageWidth())
	assert.Nil(t, err)
}

func TestGetImage100x100OK(t *testing.T) {
	filename, _ := filepath.Abs("../data/kth.jpg")
	info := &ImageInfo{filename, 100, 1.0, "jpg", true, ""}
	img, err := GetImage(info)
	assert.Equal(t, 100, img.GetImageHeight())
	assert.Equal(t, 100, img.GetImageWidth())
	assert.Nil(t, err)
}