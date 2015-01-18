package imgscale

import (
	"fmt"
	"github.com/vanng822/imgscale/imagick"
	"strings"
)

type imageProviderFile struct {
	path string
}

func (imageProvider imageProviderFile) Fetch(filename string) (*imagick.MagickWand, error) {
	img := imagick.NewMagickWand()
	err := img.ReadImage(fmt.Sprintf("%s/%s", imageProvider.path, filename))
	return img, err
}

/*
	NewImageProviderFile returns an instance of imageProviderFile
	the path is required and should point to folder where images are located
*/
func NewImageProviderFile(path string) ImageProvider {
	if path == "" {
		panic("Path can not be empty")
		// Should check if path readable here as well
	}
	provider := imageProviderFile{}
	provider.path = strings.TrimSuffix(path, "/")
	return provider
}