package imagick
/*
#include <wand/MagickWand.h>
*/
import "C"

import (
)

func GetVersion() (string, uint) {
	cnversion := C.size_t(0)
	csversion := C.MagickGetVersion(&cnversion)
	return C.GoString(csversion), uint(cnversion)
}
