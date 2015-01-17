package imgscale

import (
	"fmt"
	"github.com/gographics/imagick/imagick"
)

type Watermark struct {
	Filename string
	data      []byte
}

func (w *Watermark) load() {
	img := imagick.NewMagickWand()
	defer img.Destroy()
	if err := img.ReadImage(w.Filename); err != nil {
		panic(fmt.Sprintf("Can not load watermark '%s'", w.Filename))
	}
	w.data = img.GetImageBlob()
}

func (w *Watermark) mark(img *imagick.MagickWand) error {
	wm := imagick.NewMagickWand()
	defer wm.Destroy()
	if err := wm.ReadImageBlob(w.data); err != nil {
		return err
	}
	return img.CompositeImage(wm, imagick.COMPOSITE_OP_OVERLAY, 1, 1)
}
