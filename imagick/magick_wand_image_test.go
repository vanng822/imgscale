package imagick

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIdentifyImage(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	assert.Regexp(t, "Mime type: image", img.IdentifyImage())
}


func TestGetImageProperty(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	assert.Equal(t, "24", img.GetImageProperty("exif:Flash"))
	assert.Equal(t, "18/1, 442/100, 0/1", img.GetImageProperty("exif:GPSLongitude"))
}

func TestGetImageMimeTypeJPEG(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	assert.Equal(t, "image/jpeg", img.GetImageMimeType())
}

func TestGetImageMimeTypeGIF(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/eyes.gif")
	assert.Nil(t, err)
	assert.Equal(t, "image/gif", img.GetImageMimeType())
}

func TestGetImageMimeTypePNG(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/troll.png")
	assert.Nil(t, err)
	assert.Equal(t, "image/png", img.GetImageMimeType())
}