package imagick

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Test starting")
	Initialize()
	
	retCode := m.Run()
	
	Terminate()
	fmt.Println("Test ending")
	os.Exit(retCode)
}