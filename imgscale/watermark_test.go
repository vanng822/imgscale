package imgscale

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWatermarkLoadOK(t *testing.T) {
	watermark := Watermark{}
	watermark.Filename = "./test_data/eyes.gif"
	watermark.load()
	
	assert.Equal(t, 127, len(watermark.data))
}


func TestWatermarkLoadPanic(t *testing.T) {
	watermark := Watermark{}
	watermark.Filename = "./test_data/none.gif"
	assert.Panics(t, func() {
		watermark.load()
	})
}