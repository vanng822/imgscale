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

func (mw *MagickWand) SetImageCompressionQuality(quality uint) error {
	res := C.MagickSetImageCompressionQuality(mw.mw, C.size_t(quality))
	return mw.checkResult(BooleanType(res))
}

func (mw *MagickWand) StripImage() error {
	return mw.checkResult(BooleanType(C.MagickStripImage(mw.mw)))
}


func (mw *MagickWand) GetImageProperties(pattern string) (properties []string) {
    cspattern := C.CString(pattern)
    defer C.free(unsafe.Pointer(cspattern))
    np := C.size_t(0)
    ps := C.MagickGetImageProperties(mw.mw, cspattern, &np)
    defer relinquishMemoryCStringArray(ps)
    properties = sizedCStringArrayToStringSlice(ps, np)
    return
}

func (mw *MagickWand) GetImagePropertyValues(pattern string) map[string]string {
    properties := mw.GetImageProperties(pattern)
    if len(properties) > 0 {
        returnValues := make(map[string]string, len(properties))
        for _, property := range properties {
            returnValues[property] = mw.GetImageProperty(property)
        }
        return returnValues
    }

    return nil
}

func relinquishMemoryCStringArray(p **C.char) {
    defer C.MagickRelinquishMemory(unsafe.Pointer(p))
    for *p != nil {
        ptr := unsafe.Pointer(*p)
        if ptr != nil {
            C.MagickRelinquishMemory(ptr)
        }
        q := uintptr(unsafe.Pointer(p))
        q += unsafe.Sizeof(q)
        p = (**C.char)(unsafe.Pointer(q))
    }
}

func sizedCStringArrayToStringSlice(p **C.char, num C.size_t) []string {
    var returnStrings []string
    q := uintptr(unsafe.Pointer(p))
    for i := 0; i < int(num); i++ {
        p = (**C.char)(unsafe.Pointer(q))
        if *p == nil {
            break
        }
        returnStrings = append(returnStrings, C.GoString(*p))
        q += unsafe.Sizeof(q)
    }
    return returnStrings
}

