package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"

import (
	"unsafe"
	"regexp"
)

var (
	mimeTypePattern = regexp.MustCompile(".+Mime type: (.+)\n.*")
)

func (mw *MagickWand) IdentifyImage() string {
	return C.GoString(C.MagickIdentifyImage(mw.mw))
}

func (mw *MagickWand) GetImageProperty(property string) string {
	csProperty := C.CString(property)
	defer C.free(unsafe.Pointer(csProperty))
	csPropertyValue := C.MagickGetImageProperty(mw.mw, csProperty)
	defer mw.relinquishMemory(unsafe.Pointer(csPropertyValue))
	return C.GoString(csPropertyValue)
}

func (mw *MagickWand) GetImageMimeType() string {
	match := mimeTypePattern.FindStringSubmatch(mw.IdentifyImage())
	if len(match) > 0 {
		return match[1]
	}
	return ""
}
