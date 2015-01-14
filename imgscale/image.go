package imgscale

import (
	"github.com/gographics/imagick/imagick"
)

type ImageInfo struct {
	Filename string
	Format   *Format
	Ext      string
	Comment  string
}

type cropParams struct {
	width  uint
	height uint
	x      int
	y      int
}

func getCropParams(imageWidth, imageHeight uint, ratio float64) *cropParams {
	imageRatio := float64(imageWidth) / float64(imageHeight)
	y := 0
	x := 0
	width := imageWidth
	height := imageHeight

	if ratio < imageRatio {
		width = uint(float64(imageHeight) * ratio)
		x = int((imageWidth - width) / 2)
	} else if ratio > imageRatio {
		height = uint(float64(imageWidth) / ratio)
		y = int((imageHeight - height) / 2)
	}

	return &cropParams{width: width, height: height, x: x, y: y}
}

func scaleImage(img *imagick.MagickWand, info *ImageInfo) error {
	// no need of scaling if height is zero
	if info.Format.Height <= 0 {
		return nil
	}
	scaleFactor := float64(info.Format.Height) / float64(img.GetImageHeight())
	if info.Format.Thumbnail {
		return img.ThumbnailImage(uint(float64(img.GetImageWidth())*scaleFactor), uint(float64(img.GetImageHeight())*scaleFactor))
	} else {
		return img.ScaleImage(uint(float64(img.GetImageWidth())*scaleFactor), uint(float64(img.GetImageHeight())*scaleFactor))
	}

}

func cropImage(img *imagick.MagickWand, info *ImageInfo) error {
	params := getCropParams(img.GetImageWidth(), img.GetImageHeight(), info.Format.Ratio)
	return img.CropImage(params.width, params.height, params.x, params.y)
}

func cropScaleImage(img *imagick.MagickWand, info *ImageInfo) error {
	if err := cropImage(img, info); err != nil {
		return err
	}
	return scaleImage(img, info)
}

/*
	ProcessImage will crop/scale image to dimension specified in ImageInfo
	
*/
func ProcessImage(img *imagick.MagickWand, info *ImageInfo) (err error) {
	if info.Comment != "" {
		img.CommentImage(info.Comment)
	}
	// No crop if zero
	if info.Format.Ratio <= 0.0 {
		err = scaleImage(img, info)
	} else { // Crop first and then scale, no problem if height is zero
		err = cropScaleImage(img, info)
	}
	return err
}