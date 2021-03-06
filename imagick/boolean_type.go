package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"

type BooleanType int

const (
	BOOLEAN_TYPE_FALSE BooleanType = C.MagickFalse
	BOOLEAN_TYPE_TRUE  BooleanType = C.MagickTrue
)

func (b *BooleanType) True() bool {
	return BOOLEAN_TYPE_TRUE == *b
}

func (b *BooleanType) GoBool() bool {
	return BOOLEAN_TYPE_TRUE == *b
}