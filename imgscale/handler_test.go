package imgscale

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getHandler() *handler {
	config := &Config{}
	config.Path = "./test_data"
	config.Prefix = "img"
	config.Exts = []string{"jpg", "png", "jpeg"}
	config.Separator = "/"
	config.Formats = append(config.Formats, &Format{Prefix: "100x100", Height: 100, Ratio: 1.0, Thumbnail: true})
	config.Formats = append(config.Formats, &Format{Prefix: "original", Height: 0, Ratio: 0.0, Thumbnail: false})
	return configure(config)
}

func getHandlerDash() *handler {
	config := &Config{}
	config.Path = "./test_data"
	config.Prefix = "img"
	config.Exts = []string{"jpg", "png", "jpeg"}
	config.Separator = "-"
	config.Formats = append(config.Formats, &Format{Prefix: "100x100", Height: 100, Ratio: 1.0, Thumbnail: true})
	config.Formats = append(config.Formats, &Format{Prefix: "original", Height: 0, Ratio: 0.0, Thumbnail: false})
	return configure(config)
}

func TestHandlerMatchTrue(t *testing.T) {
	handler := getHandler()
	matched, info := handler.match("/img/original/kth.jpg")
	assert.True(t, matched)
	//assert.Equal(t, info.Ext, "jpg")
	assert.Equal(t, info.Format.Prefix, "original")

	matched2, info2 := handler.match("/img/100x100/kth.png")
	assert.True(t, matched2)
	//assert.Equal(t, info2.Ext, "png")
	assert.Equal(t, info2.Format.Prefix, "100x100")

	matched3, info3 := handler.match("/img/100x100/kth.JPEG")
	assert.True(t, matched3)
	//assert.Equal(t, info3.Ext, "jpeg")
	assert.Equal(t, info3.Format.Prefix, "100x100")
}

func TestHandlerDashMatchTrue(t *testing.T) {
	handler := getHandlerDash()
	matched, info := handler.match("/img/original-kth.jpg")
	assert.True(t, matched)
	//assert.Equal(t, info.Ext, "jpg")
	assert.Equal(t, info.Format.Prefix, "original")
	assert.Equal(t, info.Filename, "kth.jpg")

	matched2, info2 := handler.match("/img/100x100-kth.png")
	assert.True(t, matched2)
	//assert.Equal(t, info2.Ext, "png")
	assert.Equal(t, info2.Format.Prefix, "100x100")
	assert.Equal(t, info2.Filename, "kth.png")

	matched3, info3 := handler.match("/img/100x100-kt-h.png")
	assert.True(t, matched3)
	//assert.Equal(t, info3.Ext, "png")
	assert.Equal(t, info3.Format.Prefix, "100x100")
	assert.Equal(t, info3.Filename, "kt-h.png")
}

func TestHandlerMatchFalse(t *testing.T) {
	handler := getHandler()
	matched, info := handler.match("/img/kth.jpg")
	assert.False(t, matched)
	assert.Nil(t, info)

	matched2, info2 := handler.match("/img/none/kth.jpg")
	assert.False(t, matched2)
	assert.Nil(t, info2)
}

func TestHandlerDashMatchFalse(t *testing.T) {
	handler := getHandlerDash()
	matched, info := handler.match("/img/original/kth.jpg")
	assert.False(t, matched)
	assert.Nil(t, info)

	matched2, info2 := handler.match("/img/100x100/kth.png")
	assert.False(t, matched2)
	assert.Nil(t, info2)
}

func TestGetFormat(t *testing.T) {
	handler := getHandler()
	format := handler.getFormat("original")
	assert.Equal(t, format.Prefix, "original")
	assert.Equal(t, format.Height, uint(0))
	assert.False(t, format.Thumbnail)
	assert.Nil(t, handler.getFormat("tiff"))
}

func TestGetImageInfoOK(t *testing.T) {
	handler := getHandler()
	info := handler.getImageInfo("100x100", "kth.jpg")

	//assert.Equal(t, info.Ext, "jpg")
	assert.Equal(t, info.Filename, "kth.jpg")
	assert.Equal(t, info.Format.Prefix, "100x100")
	assert.Equal(t, info.Format.Height, uint(100))
	assert.Equal(t, info.Format.Ratio, 1.0)
	assert.True(t, info.Format.Thumbnail)
}

func TestGetImageInfoPanics(t *testing.T) {
	handler := getHandler()
	assert.Panics(t, func() {
		handler.getImageInfo("None", "kth.jpg")
	})
}

func TestServeHTTPOK(t *testing.T) {
	handler := getHandler()
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/img/original/kth.jpg", nil)
	assert.Nil(t, err)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 200)
	assert.InEpsilon(t, w.Body.Len(), 28611, 10)
	assert.Equal(t, w.Header().Get("Content-Type"), "image/jpeg")
}

func TestServeHTTPOKParam(t *testing.T) {
	handler := getHandler()
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/img/original/kth.jpg?test=image.png", nil)
	assert.Nil(t, err)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 200)
	assert.InEpsilon(t, w.Body.Len(), 28611, 10)
	assert.Equal(t, w.Header().Get("Content-Type"), "image/jpeg")
}

func TestServeHTTPOKPng(t *testing.T) {
	handler := getHandler()
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/img/original/troll.png", nil)
	assert.Nil(t, err)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 200)
	assert.InEpsilon(t, w.Body.Len(), 225367, 10)
	assert.Equal(t, w.Header().Get("Content-Type"), "image/png")
}

func TestServeHTTPFalseFormat(t *testing.T) {
	handler := getHandler()
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	assert.Nil(t, err)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, w.Body.Len(), 0)
}

func TestServeHTTPFalseMethod(t *testing.T) {
	handler := getHandler()
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/img/original/kth.jpg", nil)
	assert.Nil(t, err)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, w.Body.Len(), 0)
}

func TestServeHTTPFalseParams(t *testing.T) {
	handler := getHandler()
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/img/original/?t=kth.jpg", nil)
	assert.Nil(t, err)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	assert.Equal(t, w.Body.Len(), 0)
}