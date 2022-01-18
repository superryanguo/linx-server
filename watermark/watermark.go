package watermark

import (
	"fmt"
	"image"

	"github.com/noelyahan/impexp"
	"github.com/noelyahan/mergi"
)

const WaterImg = "./watermark/w1.png"

func WaterMark(base string, wm string, out string, right bool) error {
	img, err := mergi.Import(impexp.NewFileImporter(base))
	if err != nil {
		return err
	}
	wmi, err := mergi.Import(impexp.NewFileImporter(wm))
	if err != nil {
		return err
	}

	if right {
		return watermarkBottomRight(wmi, img, out)
	} else {
		return watermarkTopLeft(wmi, img, out)
	}
}

func watermarkTopLeft(wmi, img image.Image, out string) error {
	res, err := mergi.Watermark(wmi, img, image.Pt(0, 0))
	if err != nil {
		return err
	}
	//os.Remove(out) //should we remove it before we write to the same place?overwrite works?
	mergi.Export(impexp.NewFileExporter(res, out))
	return nil
}

func watermarkBottomRight(wmi, img image.Image, out string) error {
	w := wmi.Bounds()
	b := img.Bounds()
	if (b.Max.X-w.Max.X) < w.Max.X || (b.Max.Y-w.Max.Y) < w.Max.Y {
		return fmt.Errorf("Img to small for watermark!")
	}
	res, err := mergi.Watermark(wmi, img, image.Pt(b.Max.X-w.Max.X, b.Max.Y-w.Max.Y))
	if err != nil {
		return err
	}
	//os.Remove(out)
	mergi.Export(impexp.NewFileExporter(res, out))
	return nil
}
