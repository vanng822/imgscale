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

func TestDestroy(t *testing.T) {
	img := NewMagickWand()
	err := img.ReadImage("./test_data/kth.jpg")
	img.Destroy()
	assert.Nil(t, img.mw)
	assert.Nil(t, err)
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

func TestRotateImage(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	pixel := NewPixelWand()
	defer pixel.Destroy()
	err = img.RotateImage(pixel, 90)
	assert.Equal(t, int(img.GetImageWidth()), 240)
	assert.Equal(t, int(img.GetImageHeight()), 320)
	
}

func TestCompositeImage(t *testing.T) {
	img1 := NewMagickWand()
	defer img1.Destroy()
	img2 := NewMagickWand()
	defer img2.Destroy()
	img3 := NewMagickWand()
	defer img3.Destroy()
	err := img1.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	err = img2.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	err = img3.ReadImage("./test_data/eyes.gif")
	assert.Nil(t, err)
	assert.Equal(t, img1.GetImageBlob(), img2.GetImageBlob())
	err = img1.CompositeImage(img3, COMPOSITE_OP_OVERLAY, 1, 1)
	assert.Nil(t, err)
	assert.NotEqual(t, img1.GetImageBlob(), img2.GetImageBlob())
	assert.True(t, len(img1.GetImageBlob()) > len(img2.GetImageBlob()))
}

func TestImageOrientation(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	assert.Equal(t, img.GetImageOrientation(), ORIENTATION_BOTTOM_RIGHT)
	err = img.SetImageOrientation(ORIENTATION_RIGHT_BOTTOM)
	assert.Nil(t, err)
	assert.Equal(t, img.GetImageOrientation(), ORIENTATION_RIGHT_BOTTOM)
}

func TestThumbnailImage(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	err = img.ThumbnailImage(200, 200)
	assert.Nil(t, err)
	assert.Equal(t, int(img.GetImageHeight()), 200)
	assert.Equal(t, int(img.GetImageWidth()), 200)
}

func TestScaleImage(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	err = img.ScaleImage(300, 150)
	assert.Nil(t, err)
	assert.Equal(t, int(img.GetImageHeight()), 150)
	assert.Equal(t, int(img.GetImageWidth()), 300)
}

func TestCropImage(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	err = img.CropImage(100, 100, 50, 70)
	assert.Nil(t, err)
	assert.Equal(t, int(img.GetImageHeight()), 100)
	assert.Equal(t, int(img.GetImageWidth()), 100)
}

func TestCommentImage(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	err = img.CommentImage("Testing @ testing")
	assert.Nil(t, err)
}


