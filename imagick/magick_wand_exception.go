package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type MagickWandException struct {
	exceptionType        ExceptionType
	description string
}

func (mwe *MagickWandException) Error() string {
	return fmt.Sprintf("%s: %s", mwe.exceptionType.String(), mwe.description)
}
