package mongodb

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var testHost = "localhost:27017"
var testDbname = "imgscale_unittest_db"

var testUrl = fmt.Sprintf("%s/%s", testHost, testDbname)
var testPrefix = "imgscale_test_collection"
var testFilename = "kth.jpg"

func testGetImageByte() []byte {
	fd, err := os.Open("./data/kth.jpg")
	if err != nil {
		panic(err)
	}

	imageData, err := ioutil.ReadAll(fd)

	if err != nil {
		panic(err)
	}
	if len(imageData) == 0 {
		panic("No data")
	}
	return imageData
}

func testPrepareImage(filename string) {
	session := getSession(testUrl)
	defer session.Close()
	grd := session.DB(testDbname).GridFS(testPrefix)
	fd, err := grd.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	_, err = fd.Write(testGetImageByte())
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Test starting")
	session := getSession(testUrl)
	defer session.Close()
	
	retCode := m.Run()
	
	session.DB(testDbname).DropDatabase()

	fmt.Println("Test ending")
	os.Exit(retCode)
}
