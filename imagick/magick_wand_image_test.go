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
	assert.Regexp(t, "Format: JPEG", img.IdentifyImage())
}


func TestGetImageProperty(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	assert.Equal(t, "24", img.GetImageProperty("exif:Flash"))
	assert.Equal(t, "18/1, 442/100, 0/1", img.GetImageProperty("exif:GPSLongitude"))
}


func TestSetImageCompressionQualityCompare(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	assert.Nil(t, img.ReadImage("./test_data/kth.jpg"))
	img2 := NewMagickWand()
	defer img2.Destroy()
	assert.Nil(t, img2.ReadImage("./test_data/kth.jpg"))
	assert.Nil(t, img2.SetImageCompressionQuality(10))
	assert.True(t, len(img.GetImageBlob()) > len(img2.GetImageBlob()))
}

func TestStripImageCompare(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	assert.Nil(t, img.ReadImage("./test_data/kth.jpg"))
	img2 := NewMagickWand()
	defer img2.Destroy()
	assert.Nil(t, img2.ReadImage("./test_data/kth.jpg"))
	assert.Nil(t, img2.StripImage())
	assert.True(t, len(img.GetImageBlob()) > len(img2.GetImageBlob()))
}

/*
// Not working on build
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
}*/