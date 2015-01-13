package imgscale

import (
	"fmt"
)

func Configure(filename string) Handler {
	config := LoadConfig(filename)
	handler := configure(config)
	return handler
}
