package imgscale

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWatermarkLoadOK(t *testing.T) {
	watermark := Watermark{}
	watermark.Filename = "./test_data/eyes.gif"
	watermark.load()
	defer watermark.img.Destroy()
	assert.Equal(t, watermark.img.GetImageWidth(), uint(18))
	assert.Equal(t, watermark.img.GetImageHeight(), uint(18))
}


func TestWatermarkLoadPanic(t *testing.T) {
	watermark := Watermark{}
	watermark.Filename = "./test_data/none.gif"
	assert.Panics(t, func() {
		watermark.load()
	})
}