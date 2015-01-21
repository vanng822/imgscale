package imagick

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPixelWand(t *testing.T) {
	pixel := NewPixelWand()
	assert.NotNil(t, pixel)
	pixel.Destroy()
	assert.Nil(t, pixel.pw)
}