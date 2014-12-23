package imgscale

import (
	"github.com/gographics/imagick/imagick"
)

type ImageInfo struct {
	// full path to the image
	Filename string
	Height   int
	Ratio    float64
	Ext      string
	Comment  string
}

type CropParams struct {
	Width  uint
	Height uint
	X      int
	Y      int
}

func GetCropParams(imageWidth, imageHeight uint, ratio float64) *CropParams {
	/*image_ratio = (1.0 * image_width) / image_height

    # start with same size
    top = 0
    left = 0
    width = image_width
    height = image_height

    if ratio < image_ratio: # => width is larger than in the new ratio
        width = image_height * ratio
        left = (image_width - width) / 2
    elif ratio > image_ratio:
        height = image_width / ratio
        top = (image_height - height) / 2

    return CroppingParams(int(left), int(top), int(left + width), int(top + height))
    */
	imageRatio = (1.0 * imageWidth) / imageHeight
	
	return &CropParams{}
}

func scaleImage(img *imagick.MagickWand, info *ImageInfo) error {
	var scaleFactor float64
	if info.Height > 0.0 {
		scaleFactor = float64(info.Height) / float64(img.GetImageWidth())
	} else {
		scaleFactor = 1.0
	}
	return img.ScaleImage(uint(float64(img.GetImageWidth())*scaleFactor), uint(float64(img.GetImageHeight())*scaleFactor))
}

func cropImage(img *imagick.MagickWand, info *ImageInfo) error {
	params := GetCropParams(img.GetImageWidth(), img.GetImageHeight(), info.Ratio)
	return img.CropImage(params.Width, params.Height, params.X, params.Y)
}

func cropScaleImage(img *imagick.MagickWand, info *ImageInfo) error {
	if err := cropImage(img, info); err != nil {
		return err
	}
	return scaleImage(img, info)
}

func GetImage(info *ImageInfo) (*imagick.MagickWand, error) {
	img := imagick.NewMagickWand()
	err := img.ReadImage(info.Filename)

	if info.Comment != "" {
		img.CommentImage(info.Comment)
	}

	if err != nil {
		return img, err
	}

	if info.Height == 0 {
		return img, nil
	}

	if info.Ratio == 1.0 || info.Ratio == 0.0 {
		err = scaleImage(img, info)
	} else {
		//err = img.ResizeImage(uint(info.Width), uint(info.Height), imagick.FILTER_LANCZOS, 1.0)
		err = cropScaleImage(img, info)
	}
	return img, err
}
