package main

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image is a finite rectangular grid of color.Color values
// taken from a color model.

// type Image interface {
// 	// ColorModel returns the Image's color model.
// 	ColorModel() color.Model
// 	// Bounds returns the domain for which At can return non-zero color.
// 	// The bounds do not necessarily contain the point (0, 0).
// 	Bounds() Rectangle
// 	// At returns the color of the pixel at (x, y).
// 	// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// 	// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
// 	At(x, y int) color.Color
// }

type Image struct {
	w, h int
	v    uint8
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x % 256), uint8(y % 256), 255, 255}
}

func main() {
	// Rect is shorthand for Rectangle{Pt(x0, y0), Pt(x1, y1)}.
	// The returned rectangle has minimum and maximum coordinates
	// swapped if necessary so that it is well-formed.

	// NewRGBA returns a new RGBA image with the given bounds.

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())        // (0,0)-(100,100)
	fmt.Println(m.At(0, 0).RGBA()) // 0 0 0 0

	myImage := Image{200, 30, 11}
	pic.ShowImage(myImage)
}
