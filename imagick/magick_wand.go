package imagick

/*
#cgo !no_pkgconfig pkg-config: MagickWand MagickCore
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type MagickWand struct {
	mw *C.MagickWand
}

func NewMagickWand() *MagickWand {
	return &MagickWand{mw: C.NewMagickWand()}
}

func (mw *MagickWand) Destroy() {
	if mw.mw == nil {
		return
	}
	mw.mw = C.DestroyMagickWand(mw.mw)
	C.free(unsafe.Pointer(mw.mw))
	mw.mw = nil
}

func (mw *MagickWand) GetImageBlob() []byte {
	clen := C.size_t(0)
	csBlob := C.MagickGetImageBlob(mw.mw, &clen)
	defer mw.relinquishMemory(unsafe.Pointer(csBlob))
	return C.GoBytes(unsafe.Pointer(csBlob), C.int(clen))
}

func (mw *MagickWand) ReadImageBlob(blob []byte) error {
	if len(blob) == 0 {
		return fmt.Errorf("Blob can not be empty")
	}
	C.MagickReadImageBlob(mw.mw, unsafe.Pointer(&blob[0]), C.size_t(len(blob)))
	return mw.GetLastError()
}

func (mw *MagickWand) GetImageHeight() uint {
	return uint(C.MagickGetImageHeight(mw.mw))
}

func (mw *MagickWand) GetImageWidth() uint {
	return uint(C.MagickGetImageWidth(mw.mw))
}

func (mw *MagickWand) CommentImage(comment string) error {
	csComment := C.CString(comment)
	defer C.free(unsafe.Pointer(csComment))
	C.MagickCommentImage(mw.mw, csComment)
	return mw.GetLastError()
}

func (mw *MagickWand) ThumbnailImage(cols, rows uint) error {
	C.MagickThumbnailImage(mw.mw, C.size_t(cols), C.size_t(rows))
	return mw.GetLastError()
}

func (mw *MagickWand) ScaleImage(cols, rows uint) error {
	C.MagickScaleImage(mw.mw, C.size_t(cols), C.size_t(rows))
	return mw.GetLastError()
}

func (mw *MagickWand) CropImage(width, height uint, x, y int) error {
	C.MagickCropImage(mw.mw, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y))
	return mw.GetLastError()
}

func (mw *MagickWand) FlipImage() error {
	C.MagickFlipImage(mw.mw)
	return mw.GetLastError()
}

func (mw *MagickWand) RotateImage(background *PixelWand, degrees float64) error {
	C.MagickRotateImage(mw.mw, background.pw, C.double(degrees))
	return mw.GetLastError()
}

func (mw *MagickWand) FlopImage() error {
	C.MagickFlopImage(mw.mw)
	return mw.GetLastError()
}

func (mw *MagickWand) GetImageOrientation() OrientationType {
	return OrientationType(C.MagickGetImageOrientation(mw.mw))
}

func (mw *MagickWand) SetImageOrientation(orientation OrientationType) error {
	C.MagickSetImageOrientation(mw.mw, C.OrientationType(orientation))
	return mw.GetLastError()
}

func (mw *MagickWand) ReadImage(filename string) error {
	csFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csFilename))
	C.MagickReadImage(mw.mw, csFilename)
	return mw.GetLastError()
}

func (mw *MagickWand) CompositeImage(source *MagickWand, compose CompositeOperator, x, y int) error {
	C.MagickCompositeImage(mw.mw, source.mw, C.CompositeOperator(compose), C.ssize_t(x), C.ssize_t(y))
	return mw.GetLastError()
}

func (mw *MagickWand) GetLastError() error {
	var exceptionType C.ExceptionType
	csDescription := C.MagickGetException(mw.mw, &exceptionType)
	defer mw.relinquishMemory(unsafe.Pointer(csDescription))
	if ExceptionType(exceptionType) != EXCEPTION_UNDEFINED {
		mw.clearException()
		return &MagickWandException{ExceptionType(C.int(exceptionType)), C.GoString(csDescription)}
	}
	return nil
}

func (mw *MagickWand) relinquishMemory(ptr unsafe.Pointer) {
	C.MagickRelinquishMemory(ptr)
}

func (mw *MagickWand) clearException() bool {
	return 1 == C.int(C.MagickClearException(mw.mw))
}
