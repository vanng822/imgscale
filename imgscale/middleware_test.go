package imgscale

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMiddlewareConfigure(t *testing.T) {
	handler := Configure("./test_config/formats.json")
	assert.NotNil(t, handler, nil)
	assert.Implements(t, new(Handler), handler)
}

func TestMiddlewareMiddleware(t *testing.T) {
	handler := Configure("./test_config/formats.json").Middleware()
	assert.NotNil(t, handler, nil)
	assert.Implements(t, new(http.Handler), handler(nil))
}

func TestMiddlewareConfigurePanic(t *testing.T) {
	assert.Panics(t, func() {
		Configure("../test_config/formats.json")
	})
}

func TestMiddlewareConfigureWithConfig(t *testing.T) {
    conf := LoadConfig("./test_config/formats.json")
    handler := Configure(conf)
    assert.NotNil(t, handler, nil)
    assert.Implements(t, new(Handler), handler)
}
