package imgscale

import (
	"fmt"
	"github.com/gographics/imagick/imagick"
	"io/ioutil"
	"net/http"
	"strings"
)

type imageProviderHTTP struct {
	baseUrl string
}

func (imageProvider imageProviderHTTP) Fetch(filename string) (*imagick.MagickWand, error) {
	img := imagick.NewMagickWand()
	// %s%s will make it possible for using on arbitrary remote image
	// like http://127.0.0.1:8081/img/100x100/http://127.0.0.1:8080/img/original/kth.jpg
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

/*
	NewImageProviderHTTP returns an instance of imageProviderHTTP
	where baseUrl is absolute url, together with filename it should become
	a valid url to image resource.
	
	One can leave baseUrl empty and specify filename as a remote image like
	http://127.0.0.1:8081/img/100x100/http://127.0.0.1:8080/img/original/kth.jpg
	Be aware what you do and also that it may not work for some handler/frameworks
*/
func NewImageProviderHTTP(baseUrl string) ImageProvider {
	provider := imageProviderHTTP{}
	if baseUrl != "" {
		// assume we have a valid base url
		provider.baseUrl = fmt.Sprintf("%s/", strings.TrimSuffix(baseUrl, "/"))
	}
	return provider
}