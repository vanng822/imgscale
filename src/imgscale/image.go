package imgscale

import (
	"github.com/gographics/imagick/imagick"
)

func GetImage(info *ImageInfo) (*imagick.MagickWand, error) {
	img := imagick.NewMagickWand()
	err := img.ReadImage(info.Filename)

	if info.Comment != "" {
		img.CommentImage(info.Comment)
	}

	if err != nil {
		return img, err
	}

	if info.Width == 0 && info.Height == 0 {
		return img, nil
	}

	if info.KeepRatio {
		err = img.ScaleImage(uint(info.Width), uint(info.Height))
	} else {
		err = img.ResizeImage(uint(info.Width), uint(info.Height), imagick.FILTER_LANCZOS, 1.0)
	}
	return img, err
}
