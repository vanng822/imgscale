package imgscale

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchOK(t *testing.T) {
	provider := NewImageProviderFile("../data")
	img, err := provider.Fetch("kth.jpg")
	defer img.Destroy()
	assert.Nil(t, err)
	assert.Equal(t, img.GetImageWidth(), 320)
	assert.Equal(t, img.GetImageHeight(), 240)
}

func TestFetchOKSlash(t *testing.T) {
	provider := NewImageProviderFile("../data/")
	img, err := provider.Fetch("kth.jpg")
	defer img.Destroy()
	assert.Nil(t, err)
	assert.Equal(t, img.GetImageWidth(), 320)
	assert.Equal(t, img.GetImageHeight(), 240)
}

func TestFetchError(t *testing.T) {
	provider := NewImageProviderFile("../data/")
	img, err := provider.Fetch("kth2.jpg")
	defer img.Destroy()
	assert.NotNil(t, err)
}

func TestPanic(t *testing.T) {
	assert.Panics(t, func() {
		NewImageProviderFile("")
	})
}
