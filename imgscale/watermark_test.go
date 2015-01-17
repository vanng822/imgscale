package imgscale

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWatermarkLoadOK(t *testing.T) {
	watermark := Watermark{}
	watermark.Filename = "./test_data/eyes.gif"
	watermark.load()
	
	assert.Equal(t, watermark.img.GetImageWidth(), 18)
	assert.Equal(t, watermark.img.GetImageHeight(), 18)
}


func TestWatermarkLoadPanic(t *testing.T) {
	watermark := Watermark{}
	watermark.Filename = "./test_data/none.gif"
	assert.Panics(t, func() {
		watermark.load()
	})
}