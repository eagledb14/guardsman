package graphics

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewRect(w,h int, color color.Color) *ebiten.Image {
	rect := image.NewRGBA(image.Rect(0,0,w,h))

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			rect.Set(x, y, color)
		}
	}
	return ebiten.NewImageFromImage(rect)
}
