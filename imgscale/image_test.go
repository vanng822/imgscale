package imgscale

import (
	"github.com/stretchr/testify/assert"
	"github.com/vanng822/imgscale/imagick"
	"path/filepath"
	"testing"
)

func TestGetImageWrongFile(t *testing.T) {
	path, _ := filepath.Abs("./data/")
	provider := imageProviderFile{path}
	
	filename := "kth.jpg"
	f := &Format{Prefix: "100x100", Height: 100, Ratio: 1.0, Thumbnail: true}
	info := &ImageInfo{filename, f, "jpg", ""}
	_, err := provider.Fetch(info.Filename)
	assert.Error(t, err)
}

func TestGetImageScaleOK(t *testing.T) {
	path, _ := filepath.Abs("./test_data/")
	provider := imageProviderFile{path}
	filename := "kth.jpg"
	f := &Format{Prefix: "133x100", Height: 100, Ratio: 0.0, Thumbnail: false}
	info := &ImageInfo{filename, f, "jpg", ""}
	img, err := provider.Fetch(info.Filename)
	assert.Nil(t, err)
	err = ProcessImage(img, info)
	assert.Equal(t, 100, img.GetImageHeight())
	assert.Equal(t, 133, img.GetImageWidth())
	assert.Nil(t, err)
}

func TestGetImage100x100OK(t *testing.T) {
	path, _ := filepath.Abs("./test_data/")
	provider := NewImageProviderFile(path)
	filename := "kth.jpg"
	
	f := &Format{Prefix: "100x100", Height: 100, Ratio: 1.0, Thumbnail: true}
	info := &ImageInfo{filename, f, "jpg", ""}
	img, err := provider.Fetch(info.Filename)
	assert.Nil(t, err)
	err = ProcessImage(img, info)
	assert.Equal(t, 100, img.GetImageHeight())
	assert.Equal(t, 100, img.GetImageWidth())
	assert.Nil(t, err)
}

func TestAutoRotate(t *testing.T) {
	path, _ := filepath.Abs("./test_data/")
	provider := NewImageProviderFile(path)
	filename := "kth.jpg"
	
	f := &Format{Prefix: "100x75", Height: 100, Ratio: 1.335, Thumbnail: true}
	info := &ImageInfo{filename, f, "jpg", ""}
	img, err := provider.Fetch(info.Filename)
	assert.Nil(t, err)
	rotateOrientations := make([]imagick.OrientationType, 0)
	rotateOrientations = append(rotateOrientations, imagick.ORIENTATION_TOP_RIGHT)
	rotateOrientations = append(rotateOrientations, imagick.ORIENTATION_BOTTOM_RIGHT)
	rotateOrientations = append(rotateOrientations, imagick.ORIENTATION_BOTTOM_LEFT)
	rotateOrientations = append(rotateOrientations, imagick.ORIENTATION_RIGHT_TOP)
	rotateOrientations = append(rotateOrientations, imagick.ORIENTATION_LEFT_BOTTOM)
	rotateOrientations = append(rotateOrientations, imagick.ORIENTATION_TOP_LEFT)
	
	for _, orientation := range rotateOrientations {
		img.SetImageOrientation(orientation)
		err = AutoRotate(img)
		assert.Nil(t, err)
		assert.Equal(t, imagick.ORIENTATION_TOP_LEFT, img.GetImageOrientation())
	}
	
	noRotate := make([]imagick.OrientationType, 0)
	noRotate = append(noRotate, imagick.ORIENTATION_LEFT_TOP)
	noRotate = append(noRotate, imagick.ORIENTATION_UNDEFINED)
	noRotate = append(noRotate, imagick.ORIENTATION_RIGHT_BOTTOM)
	noRotate = append(noRotate, imagick.ORIENTATION_LEFT_TOP)
	
	for _, orientation := range noRotate {
		img.SetImageOrientation(orientation)
		err = AutoRotate(img)
		assert.Nil(t, err)
		assert.Equal(t, orientation, img.GetImageOrientation())
	}
	
}
