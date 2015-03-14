package http

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpFetchOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./data/kth.jpg")
	}))
	defer ts.Close()

	provider := New("")
	img, err := provider.Fetch(ts.URL)
	defer img.Destroy()

	assert.Nil(t, err)
	assert.Equal(t, img.GetImageWidth(), uint(320))
	assert.Equal(t, img.GetImageHeight(), uint(240))
}

func TestHttpFetchOKBaseUrl(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./data/kth.jpg")
	}))
	defer ts.Close()

	provider := New(ts.URL)
	img, err := provider.Fetch("kth.jpg")
	defer img.Destroy()

	assert.Nil(t, err)
	assert.Equal(t, img.GetImageWidth(), uint(320))
	assert.Equal(t, img.GetImageHeight(), uint(240))
}
