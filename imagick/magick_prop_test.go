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