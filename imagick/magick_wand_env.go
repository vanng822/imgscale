package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"

var (
	initialized bool
)

func Initialize() {
	if initialized {
		return
	}
	C.MagickWandGenesis()
	initialized = true
}

func Terminate() {
	if initialized {
		C.MagickWandTerminus()
		initialized = false
	}
}
