package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 20)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{55, 94, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
