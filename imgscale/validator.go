package imgscale

import (
	"strings"
)

type defaultValidator struct {}

func (v defaultValidator) Validate(filename string) bool {
	return strings.Index(filename, "..") == -1
}
