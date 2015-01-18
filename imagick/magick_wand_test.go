package imagick

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestReadImageWrongFile(t *testing.T) {
	filename, _ := filepath.Abs("./test_data/none.jpg")
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage(filename)

	assert.Error(t, err)
	assert.Regexp(t, "BlobError: unable to open image", err.Error())
}

func TestReadImageOK(t *testing.T) {
	filename, _ := filepath.Abs("./test_data/kth.jpg")

	img := NewMagickWand()
	defer img.Destroy()
	
	err := img.ReadImage(filename)
	assert.Nil(t, err)
	assert.Equal(t, int(img.GetImageWidth()), 320)
	assert.Equal(t, int(img.GetImageHeight()), 240)
}