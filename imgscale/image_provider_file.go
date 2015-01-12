package imgscale

import (
	"fmt"
	"github.com/gographics/imagick/imagick"
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