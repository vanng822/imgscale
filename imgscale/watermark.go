package imgscale

import (
	"fmt"
	"github.com/gographics/imagick/imagick"
)

type Watermark struct {
	Filename string
	img      *imagick.MagickWand
}

func (w *Watermark) load() {
	w.img = imagick.NewMagickWand()
	if err := w.img.ReadImage(w.Filename); err != nil {
		panic(fmt.Sprintf("Can not load watermark '%s'", w.Filename))
	}
}

func (w *Watermark) mark(img *imagick.MagickWand) error {
	return img.CompositeImage(w.img, imagick.COMPOSITE_OP_OVERLAY, 1, 1)
}
