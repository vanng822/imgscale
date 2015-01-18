package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type ExceptionType int

const (
	EXCEPTION_UNDEFINED          ExceptionType = C.UndefinedException
	EXCEPTION_WARNING            ExceptionType = C.WarningException
	WARNING_RESOURCE_LIMIT       ExceptionType = C.ResourceLimitWarning
	WARNING_TYPE                 ExceptionType = C.TypeWarning
	WARNING_OPTION               ExceptionType = C.OptionWarning
	WARNING_DELEGATE             ExceptionType = C.DelegateWarning
	WARNING_MISSING_DELEGATE     ExceptionType = C.MissingDelegateWarning
	WARNING_CORRUPT_IMAGE        ExceptionType = C.CorruptImageWarning
	WARNING_FILE_OPEN            ExceptionType = C.FileOpenWarning
	WARNING_BLOB                 ExceptionType = C.BlobWarning
	WARNING_STREAM               ExceptionType = C.StreamWarning
	WARNING_CACHE                ExceptionType = C.CacheWarning
	WARNING_CODER                ExceptionType = C.CoderWarning
	WARNING_FILTER               ExceptionType = C.FilterWarning
	WARNING_MODULE               ExceptionType = C.ModuleWarning
	WARNING_DRAW                 ExceptionType = C.DrawWarning
	WARNING_IMAGE                ExceptionType = C.ImageWarning
	WARNING_WAND                 ExceptionType = C.WandWarning
	WARNING_RANDOM               ExceptionType = C.RandomWarning
	WARNING_XSERVER              ExceptionType = C.XServerWarning
	WARNING_MONITOR              ExceptionType = C.MonitorWarning
	WARNING_REGISTRY             ExceptionType = C.RegistryWarning
	WARNING_CONFIGURE            ExceptionType = C.ConfigureWarning
	WARNING_POLICY               ExceptionType = C.PolicyWarning
	EXCEPTION_ERROR              ExceptionType = C.ErrorException
	ERROR_RESOURCE_LIMIT         ExceptionType = C.ResourceLimitError
	ERROR_TYPE                   ExceptionType = C.TypeError
	ERROR_OPTION                 ExceptionType = C.OptionError
	ERROR_DELEGATE               ExceptionType = C.DelegateError
	ERROR_MISSING_DELEGATE       ExceptionType = C.MissingDelegateError
	ERROR_CORRUPT_IMAGE          ExceptionType = C.CorruptImageError
	ERROR_FILE_OPEN              ExceptionType = C.FileOpenError
	ERROR_BLOB                   ExceptionType = C.BlobError
	ERROR_STREAM                 ExceptionType = C.StreamError
	ERROR_CACHE                  ExceptionType = C.CacheError
	ERROR_CODER                  ExceptionType = C.CoderError
	ERROR_FILTER                 ExceptionType = C.FilterError
	ERROR_MODULE                 ExceptionType = C.ModuleError
	ERROR_DRAW                   ExceptionType = C.DrawError
	ERROR_IMAGE                  ExceptionType = C.ImageError
	ERROR_WAND                   ExceptionType = C.WandError
	ERROR_RANDOM                 ExceptionType = C.RandomError
	ERROR_XSERVER                ExceptionType = C.XServerError
	ERROR_MONITOR                ExceptionType = C.MonitorError
	ERROR_REGISTRY               ExceptionType = C.RegistryError
	ERROR_CONFIGURE              ExceptionType = C.ConfigureError
	ERROR_POLICY                 ExceptionType = C.PolicyError
	EXCEPTION_FATAL_ERROR        ExceptionType = C.FatalErrorException
	FATAL_ERROR_RESOURCE_LIMIT   ExceptionType = C.ResourceLimitFatalError
	FATAL_ERROR_TYPE             ExceptionType = C.TypeFatalError
	FATAL_ERROR_OPTION           ExceptionType = C.OptionFatalError
	FATAL_ERROR_DELEGATE         ExceptionType = C.DelegateFatalError
	FATAL_ERROR_MISSING_DELEGATE ExceptionType = C.MissingDelegateFatalError
	FATAL_ERROR_CORRUPT_IMAGE    ExceptionType = C.CorruptImageFatalError
	FATAL_ERROR_FILE_OPEN        ExceptionType = C.FileOpenFatalError
	FATAL_ERROR_BLOB             ExceptionType = C.BlobFatalError
	FATAL_ERROR_STREAM           ExceptionType = C.StreamFatalError
	FATAL_ERROR_CACHE            ExceptionType = C.CacheFatalError
	FATAL_ERROR_CODER            ExceptionType = C.CoderFatalError
	FATAL_ERROR_FILTER           ExceptionType = C.FilterFatalError
	FATAL_ERROR_MODULE           ExceptionType = C.ModuleFatalError
	FATAL_ERROR_DRAW             ExceptionType = C.DrawFatalError
	FATAL_ERROR_IMAGE            ExceptionType = C.ImageFatalError
	FATAL_ERROR_WAND             ExceptionType = C.WandFatalError
	FATAL_ERROR_RANDOM           ExceptionType = C.RandomFatalError
	FATAL_ERROR_XSERVER          ExceptionType = C.XServerFatalError
	FATAL_ERROR_MONITOR          ExceptionType = C.MonitorFatalError
	FATAL_ERROR_REGISTRY         ExceptionType = C.RegistryFatalError
	FATAL_ERROR_CONFIGURE        ExceptionType = C.ConfigureFatalError
	FATAL_ERROR_POLICY           ExceptionType = C.PolicyFatalError
)

var exceptionTypeString = map[ExceptionType]string{
	EXCEPTION_UNDEFINED:          "UndefinedException",
	EXCEPTION_WARNING:            "WarningException",
	//WARNING_RESOURCE_LIMIT:       "ResourceLimitWarning",
	WARNING_TYPE:                 "TypeWarning",
	WARNING_OPTION:               "OptionWarning",
	WARNING_DELEGATE:             "DelegateWarning",
	WARNING_MISSING_DELEGATE:     "MissingDelegateWarning",
	WARNING_CORRUPT_IMAGE:        "CorruptImageWarning",
	WARNING_FILE_OPEN:            "FileOpenWarning",
	WARNING_BLOB:                 "BlobWarning",
	WARNING_STREAM:               "StreamWarning",
	WARNING_CACHE:                "CacheWarning",
	WARNING_CODER:                "CoderWarning",
	WARNING_FILTER:               "FilterWarning",
	WARNING_MODULE:               "ModuleWarning",
	WARNING_DRAW:                 "DrawWarning",
	WARNING_IMAGE:                "ImageWarning",
	WARNING_WAND:                 "WandWarning",
	WARNING_RANDOM:               "RandomWarning",
	WARNING_XSERVER:              "XServerWarning",
	WARNING_MONITOR:              "MonitorWarning",
	WARNING_REGISTRY:             "RegistryWarning",
	WARNING_CONFIGURE:            "ConfigureWarning",
	WARNING_POLICY:               "PolicyWarning",
	EXCEPTION_ERROR:              "ErrorException",
	//ERROR_RESOURCE_LIMIT:         "ResourceLimitError",
	ERROR_TYPE:                   "TypeError",
	ERROR_OPTION:                 "OptionError",
	ERROR_DELEGATE:               "DelegateError",
	ERROR_MISSING_DELEGATE:       "MissingDelegateError",
	ERROR_CORRUPT_IMAGE:          "CorruptImageError",
	ERROR_FILE_OPEN:              "FileOpenError",
	ERROR_BLOB:                   "BlobError",
	ERROR_STREAM:                 "StreamError",
	ERROR_CACHE:                  "CacheError",
	ERROR_CODER:                  "CoderError",
	ERROR_FILTER:                 "FilterError",
	ERROR_MODULE:                 "ModuleError",
	ERROR_DRAW:                   "DrawError",
	ERROR_IMAGE:                  "ImageError",
	ERROR_WAND:                   "WandError",
	ERROR_RANDOM:                 "RandomError",
	ERROR_XSERVER:                "XServerError",
	ERROR_MONITOR:                "MonitorError",
	ERROR_REGISTRY:               "RegistryError",
	ERROR_CONFIGURE:              "ConfigureError",
	ERROR_POLICY:                 "PolicyError",
	EXCEPTION_FATAL_ERROR:        "FatalErrorException",
	//FATAL_ERROR_RESOURCE_LIMIT:   "ResourceLimitFatalError",
	FATAL_ERROR_TYPE:             "TypeFatalError",
	FATAL_ERROR_OPTION:           "OptionFatalError",
	FATAL_ERROR_DELEGATE:         "DelegateFatalError",
	FATAL_ERROR_MISSING_DELEGATE: "MissingDelegateFatalError",
	FATAL_ERROR_CORRUPT_IMAGE:    "CorruptImageFatalError",
	FATAL_ERROR_FILE_OPEN:        "FileOpenFatalError",
	FATAL_ERROR_BLOB:             "BlobFatalError",
	FATAL_ERROR_STREAM:           "StreamFatalError",
	FATAL_ERROR_CACHE:            "CacheFatalError",
	FATAL_ERROR_CODER:            "CoderFatalError",
	FATAL_ERROR_FILTER:           "FilterFatalError",
	FATAL_ERROR_MODULE:           "ModuleFatalError",
	FATAL_ERROR_DRAW:             "DrawFatalError",
	FATAL_ERROR_IMAGE:            "ImageFatalError",
	FATAL_ERROR_WAND:             "WandFatalError",
	FATAL_ERROR_RANDOM:           "RandomFatalError",
	FATAL_ERROR_XSERVER:          "XServerFatalError",
	FATAL_ERROR_MONITOR:          "MonitorFatalError",
	FATAL_ERROR_REGISTRY:         "RegistryFatalError",
	FATAL_ERROR_CONFIGURE:        "ConfigureFatalError",
	FATAL_ERROR_POLICY:           "PolicyFatalError",
}

func (exceptionType *ExceptionType) String() string {
	if v, ok := exceptionTypeString[ExceptionType(*exceptionType)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownError[%d]", *exceptionType)
}
