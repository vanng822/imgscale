package imgscale

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiddlewareConfigure(t *testing.T) {
	handler := Configure("./test_config/formats.json")
	assert.NotNil(t, handler, nil)
	assert.Implements(t, new(Handler), handler)
}
