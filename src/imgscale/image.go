package imgscale

import (
	"github.com/gographics/imagick/imagick"
)

func GetImage(info *ImageInfo) *imagick.MagickWand {
	img := imagick.NewMagickWand()
	err := img.ReadImage(info.Filename)
	if err != nil {
		panic(err)
	}

	if info.Original {
		return img
	}

	if info.KeepRatio {
		err = img.ScaleImage(uint(info.Width), uint(info.Height))
		if err != nil {
			panic(err)
		}
	} else {
		err = img.ResizeImage(uint(info.Width), uint(info.Height), imagick.FILTER_LANCZOS, 1.0)
		if err != nil {
			panic(err)
		}
	}
	return img
}
