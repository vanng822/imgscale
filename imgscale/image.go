package imgscale

import (
	"fmt"
	"github.com/gographics/imagick/imagick"
	"io/ioutil"
	"net/http"
	"strings"
)

type ImageInfo struct {
	Filename string
	Format   *Format
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

	return &CropParams{Width: width, Height: height, X: x, Y: y}
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
	params := GetCropParams(img.GetImageWidth(), img.GetImageHeight(), info.Format.Ratio)
	return img.CropImage(params.Width, params.Height, params.X, params.Y)
}

func cropScaleImage(img *imagick.MagickWand, info *ImageInfo) error {
	if err := cropImage(img, info); err != nil {
		return err
	}
	return scaleImage(img, info)
}

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