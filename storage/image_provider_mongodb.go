package storage

import (
	"github.com/vanng822/imgscale/imagick"
	"github.com/vanng822/imgscale/imgscale"
	"gopkg.in/mgo.v2"
	"io/ioutil"
)

var session *mgo.Session

func dial(url string) *mgo.Session {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	return session
}

func getSession(url string) *mgo.Session {
	if session == nil {
		session = dial(url)
	}
	return session.New()
}

func getDB(url string) *mgo.Database {
	return getSession(url).DB("")
}

type imageProviderMongodb struct {
	url    string
	prefix string
}

func NewImageProviderMongodb(prefix, url string) imgscale.ImageProvider {
	return &imageProviderMongodb{
		url:    url,
		prefix: prefix,
	}
}

func (ipm *imageProviderMongodb) getGridFS() *mgo.GridFS {
	return getDB(ipm.url).GridFS(ipm.prefix)
}

func (ipm *imageProviderMongodb) Fetch(filename string) (*imagick.MagickWand, error) {
	img := imagick.NewMagickWand()
	gridfs := ipm.getGridFS()

	fd, err := gridfs.Open(filename)
	if err != nil {
		return img, err
	}
	defer fd.Close()
	imageData, err := ioutil.ReadAll(fd)
	if err != nil {
		return img, err
	}
	err = img.ReadImageBlob(imageData)
	return img, err
}
