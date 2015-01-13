package imgscale

import (
)

func Configure(filename string) Handler {
	config := LoadConfig(filename)
	return configure(config)
}
