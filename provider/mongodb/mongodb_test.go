package mongodb

import (
	"github.com/stretchr/testify/assert"
	"github.com/vanng822/imgscale/imgscale"
	"testing"
)

func TestInterfaceImplements(t *testing.T) {
	s := New(testUrl, testPrefix)
	assert.Implements(t, new(imgscale.ImageProvider), s)
}

func TestFetch(t *testing.T) {
	s := imageProviderMongodb{
		url:    testUrl,
		prefix: testPrefix,
	}
	testPrepareImage(testFilename)
	
	img, err := s.Fetch(testFilename)
	
	assert.Nil(t, err)
	assert.NotNil(t, img)
	assert.Equal(t, img.GetImageWidth(), uint(320))
	assert.Equal(t, img.GetImageHeight(), uint(240))
}