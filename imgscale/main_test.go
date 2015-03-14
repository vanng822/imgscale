package imgscale

import (
	"fmt"
	"github.com/vanng822/imgscale/imagick"
	"os"
	"io/ioutil"
	"testing"
)


func testGetImageByte(filename string) []byte {
	fd, err := os.Open(filename)
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

func TestMain(m *testing.M) {
	fmt.Println("Test starting")
	imagick.Initialize()

	retCode := m.Run()
	
	
	imagick.Terminate()
	fmt.Println("Test ending")
	os.Exit(retCode)
}
