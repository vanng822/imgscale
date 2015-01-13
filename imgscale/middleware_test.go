package imgscale

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiddlewareConfigure(t *testing.T) {
	handler := Configure("../config/formats.json")
	assert.NotNil(t, handler, nil)
	assert.Implements(t, new(Handler), handler)
}
