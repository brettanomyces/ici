package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width, height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel

}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	red := uint8(0)
	green := uint8(100)
	blue := uint8(200)
	alpha := uint8(255)
	return color.RGBA{red, green, blue, alpha}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
