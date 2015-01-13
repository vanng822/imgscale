package imgscale

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestHttpFetchOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../data/kth.jpg")
	}))
	defer ts.Close()
	
	provider := NewImageProviderHTTP("")
	img, err := provider.Fetch(ts.URL)
	defer img.Destroy()
	
	assert.Nil(t, err)
	assert.Equal(t, img.GetImageWidth(), 320)
	assert.Equal(t, img.GetImageHeight(), 240)
}