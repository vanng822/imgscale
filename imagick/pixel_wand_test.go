package imagick

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPixelWand(t *testing.T) {
	pixel := NewPixelWand()
	defer pixel.Destroy()
	assert.NotNil(t, pixel)
}