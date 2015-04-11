package s3

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"github.com/vanng822/imgscale/imagick"
	"github.com/vanng822/imgscale/imgscale"
)

type imageProviderS3 struct {
	bucket *s3.Bucket
	ACL    s3.ACL
}

func New(config map[string]string) imgscale.ImageProvider {
	regionName := config["regionName"]
	accessKey := config["accessKey"]
	secretKey := config["secretKey"]
	bucketName := config["bucketName"]
	if regionName == "" || accessKey == "" || secretKey == "" || bucketName == "" {
		panic("Config must contain regionName, accessKey, secretKey and bucketName")
	}

	ACL := config["ACL"]

	auth, err := aws.GetAuth(accessKey, secretKey)
	if err != nil {
		panic("Could not get auth")
	}

	client := s3.New(auth, aws.Regions[regionName])
	bucket := client.Bucket(bucketName)

	return NewIPS3(bucket, s3.ACL(ACL))
}

func NewIPS3(bucket *s3.Bucket, ACL s3.ACL) imgscale.ImageProvider {
	s := &imageProviderS3{
		bucket: bucket,
		ACL:    ACL,
	}

	return s
}

func (ips3 *imageProviderS3) Fetch(filename string) (*imagick.MagickWand, error) {
	imgData, err := ips3.bucket.Get(filename)
	img := imagick.NewMagickWand()
	
	if err != nil {
		return img, err
	}
	
	err = img.ReadImageBlob(imgData)
	return img, err
}

