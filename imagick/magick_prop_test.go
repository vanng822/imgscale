package imagick

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVersion(t *testing.T) {
	sversion, nversion := GetVersion()
	assert.Regexp(t, "ImageMagick", sversion)
	assert.True(t, nversion > 0)
}


func TestResourceLimit(t *testing.T) {
	area := GetResourceLimit(RESOURCE_AREA)/2
	assert.True(t, SetResourceLimit(RESOURCE_AREA, area))
	assert.Equal(t, area, GetResourceLimit(RESOURCE_AREA))
}