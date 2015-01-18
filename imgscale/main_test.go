package imgscale

import (
	"fmt"
	"github.com/vanng822/imgscale/imagick"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Test starting")
	imagick.Initialize()

	retCode := m.Run()
	
	
	imagick.Terminate()
	fmt.Println("Test ending")
	os.Exit(retCode)
}
