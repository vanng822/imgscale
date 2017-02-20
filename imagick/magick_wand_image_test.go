package imagick

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentifyImage(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	assert.Regexp(t, "Format: JPEG", img.IdentifyImage())
}

func TestGetImageProperty(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	assert.Equal(t, "24", img.GetImageProperty("exif:Flash"))
	assert.Equal(t, "18/1, 442/100, 0/1", img.GetImageProperty("exif:GPSLongitude"))
}

func TestGetImagePropertyValues(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	expected := map[string]string{"exif:GPSAltitude": "83329/2688", "exif:GPSLatitude": "59/1, 2082/100, 0/1", "exif:SubjectArea": "1295, 967, 699, 696", "exif:ExposureProgram": "2", "exif:Flash": "24", "exif:GPSInfo": "594", "exif:MeteringMode": "5", "exif:Model": "iPhone 4", "exif:YCbCrPositioning": "1", "jpeg:sampling-factor": "2x2,1x1,1x1", "exif:ExifOffset": "204", "exif:ExifVersion": "48, 50, 50, 49", "exif:ComponentsConfiguration": "0, 0, 0, 1", "exif:ExifImageLength": "1936", "exif:ExifImageWidth": "2592", "exif:GPSLongitudeRef": "E", "exif:GPSTimeStamp": "17/1, 20/1, 1319/100", "jpeg:colorspace": "2", "exif:BrightnessValue": "6769/958", "exif:ColorSpace": "1", "exif:ISOSpeedRatings": "80", "exif:Make": "Apple", "exif:Orientation": "3", "exif:SceneCaptureType": "0", "exif:Software": "6.1.3", "exif:WhiteBalance": "0", "exif:GPSAltitudeRef": "0", "exif:ExposureTime": "1/274", "exif:FlashPixVersion": "48, 49, 48, 48", "exif:YResolution": "72/1", "exif:ApertureValue": "4281/1441", "exif:ExposureMode": "0", "exif:GPSLatitudeRef": "N", "exif:GPSLongitude": "18/1, 442/100, 0/1", "exif:ShutterSpeedValue": "6285/776", "exif:FNumber": "14/5", "exif:GPSImgDirection": "24452/637", "exif:GPSImgDirectionRef": "T", "exif:ResolutionUnit": "2", "exif:SensingMethod": "2", "exif:FocalLength": "77/20", "exif:FocalLengthIn35mmFilm": "35", "exif:DateTimeOriginal": "2013:05:07 19:20:18", "exif:XResolution": "72/1", "exif:DateTime": "2013:05:07 19:20:18", "exif:DateTimeDigitized": "2013:05:07 19:20:18"}
	result := img.GetImagePropertyValues("*")
	for key, value := range expected {
		rvalue, ok := result[key]
		assert.True(t, ok)
		assert.Equal(t, value, rvalue)
	}

}

func TestSetImageCompressionQualityCompare(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	assert.Nil(t, img.ReadImage("./test_data/kth.jpg"))
	img2 := NewMagickWand()
	defer img2.Destroy()
	assert.Nil(t, img2.ReadImage("./test_data/kth.jpg"))
	assert.Nil(t, img2.SetImageCompressionQuality(10))
	assert.True(t, len(img.GetImageBlob()) > len(img2.GetImageBlob()))
}

func TestStripImageCompare(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	assert.Nil(t, img.ReadImage("./test_data/kth.jpg"))
	img2 := NewMagickWand()
	defer img2.Destroy()
	assert.Nil(t, img2.ReadImage("./test_data/kth.jpg"))
	assert.Nil(t, img2.StripImage())
	assert.True(t, len(img.GetImageBlob()) > len(img2.GetImageBlob()))
}

/*
// Not working on build
func TestGetImageMimeTypeJPEG(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/kth.jpg")
	assert.Nil(t, err)
	assert.Equal(t, "image/jpeg", img.GetImageMimeType())
}

func TestGetImageMimeTypeGIF(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/eyes.gif")
	assert.Nil(t, err)
	assert.Equal(t, "image/gif", img.GetImageMimeType())
}

func TestGetImageMimeTypePNG(t *testing.T) {
	img := NewMagickWand()
	defer img.Destroy()
	err := img.ReadImage("./test_data/troll.png")
	assert.Nil(t, err)
	assert.Equal(t, "image/png", img.GetImageMimeType())
}*/
