package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"

type MagickBooleanType int

const(
	BOOLEAN_TYPE_FALSE MagickBooleanType = C.MagickFalse
	BOOLEAN_TYPE_TRUE MagickBooleanType = C.MagickTrue
)