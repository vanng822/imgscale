package imgscale

import (
	"fmt"
	"github.com/gographics/imagick/imagick"
	"io/ioutil"
	"net/http"
	"strings"
)

type ImageProviderFile struct {
	Path string
}

func (imageProvider ImageProviderFile) Fetch(info *ImageInfo) (*imagick.MagickWand, error) {
	img := imagick.NewMagickWand()
	err := img.ReadImage(fmt.Sprintf("%s/%s", imageProvider.Path, info.Filename))
	return img, err
}

func NewImageProviderFile(path string) ImageProvider {
	if path == "" {
		panic("Path can not be empty")
		// Should check if path readable here as well
	}
	provider := ImageProviderFile{}
	provider.Path = strings.TrimSuffix(path, "/")
	return provider
}

type ImageProviderHTTP struct {
	BaseUrl string
}

func (imageProvider ImageProviderHTTP) Fetch(info *ImageInfo) (*imagick.MagickWand, error) {
	img := imagick.NewMagickWand()
	// %s%s will make it possible for using on arbitrary remote image
	// like http://127.0.0.1:8081/img/100x100/http://127.0.0.1:8080/img/original/kth.jpg
	resp, err := http.Get(fmt.Sprintf("%s%s", imageProvider.BaseUrl, info.Filename))
	if err != nil {
		return img, err
	}
	defer resp.Body.Close()
	
	imgData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return img, err
	}
	err = img.ReadImageBlob(imgData)
	return img, err
}

func NewImageProviderHTTP(baseUrl string) ImageProvider {
	provider := ImageProviderHTTP{}
	if baseUrl != "" {
		// assume we have a valid base url
		provider.BaseUrl = fmt.Sprintf("%s/", strings.TrimSuffix(baseUrl, "/"))
	}
	return provider
}