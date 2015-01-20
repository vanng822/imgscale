package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"

import ()

func GetVersion() (string, uint) {
	cnversion := C.size_t(0)
	csversion := C.MagickGetVersion(&cnversion)
	return C.GoString(csversion), uint(cnversion)
}

func SetResourceLimit(resourceType ResourceType, limit int64) bool {
	res := C.MagickSetResourceLimit(C.ResourceType(resourceType), C.MagickSizeType(limit))
	return BOOLEAN_TYPE_TRUE == BooleanType(res)

}

// Returns the specified resource limit in megabytes.
func GetResourceLimit(resourceType ResourceType) int64 {
	return int64(C.MagickGetResourceLimit(C.ResourceType(resourceType)))
}
