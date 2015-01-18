package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"

import (
	"unsafe"
)

type PixelWand struct {
	pw *C.PixelWand
}

func NewPixelWand() *PixelWand {
	return &PixelWand{pw: C.NewPixelWand()}
}

func (pw *PixelWand) Destroy() {
	if pw.pw == nil {
		return
	}
	pw.pw = C.DestroyPixelWand(pw.pw)
	C.free(unsafe.Pointer(pw.pw))
	pw.pw = nil
}
