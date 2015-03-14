# Image Provider

This is where you can find image provider for imgscale. Each provider should has own package, like http and mongodb. Bellow is the interface to implement and example.

## Interface

	type ImageProvider interface {
		Fetch(filename string) (*imagick.MagickWand, error)
	}
	
## Example
	
	package http

	import (
		"fmt"
		"github.com/vanng822/imgscale/imagick"
		"github.com/vanng822/imgscale/imgscale"
		"io/ioutil"
		"net/http"
		"strings"
	)
	
	type imageProviderHTTP struct {
		baseUrl string
	}
	
	func (imageProvider *imageProviderHTTP) Fetch(filename string) (*imagick.MagickWand, error) {
		img := imagick.NewMagickWand()
		resp, err := http.Get(fmt.Sprintf("%s%s", imageProvider.baseUrl, filename))
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
	
	func New(baseUrl string) imgscale.ImageProvider {
		provider := &imageProviderHTTP{}
		if baseUrl != "" {
			provider.baseUrl = fmt.Sprintf("%s/", strings.TrimSuffix(baseUrl, "/"))
		}
		return provider
	}