package main

import (
	"fmt"
	"image"
	"log"

	"github.com/noelyahan/impexp"
	"github.com/noelyahan/mergi"
)

func main() {
	img, err := mergi.Import(impexp.NewFileImporter("test1.jpg"))
	if err != nil {
		log.Fatal(err)
	}
	watermarkImage, err := mergi.Import(impexp.NewFileImporter("w1.png"))
	if err != nil {
		log.Fatal(err)
	}

	watermarkBottomRight(watermarkImage, img)
}

func watermarkTopLeft(watermarkImage, img image.Image) {
	res, err := mergi.Watermark(watermarkImage, img, image.Pt(0, 0))
	if err != nil {
		log.Fatal(err)
	}
	mergi.Export(impexp.NewFileExporter(res, "./mywater2.png"))
}
func watermarkBottomRight(watermarkImage, img image.Image) {
	w := watermarkImage.Bounds()
	b := img.Bounds()
	if (b.Max.X-w.Max.X) < w.Max.X || (b.Max.Y-w.Max.Y) < w.Max.Y {
		log.Fatal(fmt.Errorf("Img to small for watermark!"))
	}
	res, err := mergi.Watermark(watermarkImage, img, image.Pt(b.Max.X-w.Max.X, b.Max.Y-w.Max.Y))
	if err != nil {
		log.Fatal(err)
	}
	mergi.Export(impexp.NewFileExporter(res, "./mywater2.png"))
}
