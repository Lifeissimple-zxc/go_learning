package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w int
	h int
	v int
}

// our Image needs to implement and interface: https://pkg.go.dev/image#Image

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(i.v), uint8(i.v / 2), 255, 255}
}

func main() {
	m := Image{
		w: 10,
		h: 10,
		v: 0,
	}
	pic.ShowImage(m)
}
