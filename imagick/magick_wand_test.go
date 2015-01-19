package imagick

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
	"os"
	"io/ioutil"
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

func TestGetImageBlob(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	assert.InEpsilon(t, len(img.GetImageBlob()), 28611, 10)
}


func TestReadImageBlob(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	
	file, err := os.Open("./test_data/kth.jpg")
	assert.Nil(t, err)
	defer file.Close()
	imgData, err := ioutil.ReadAll(file)
	assert.Nil(t, err)
	err = img.ReadImageBlob(imgData)
	assert.Nil(t, err)
	assert.InEpsilon(t, len(img.GetImageBlob()), 28611, 10)
	assert.Equal(t, int(img.GetImageWidth()), 320)
	assert.Equal(t, int(img.GetImageHeight()), 240)
}