package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	img *ebiten.Image
	Op *ebiten.DrawImageOptions
	w, h int
	Offset FloatPoint
	multiplier int
}

func NewCamera(w, h int) *Camera {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w), -float64(h))

	img := ebiten.NewImage(w * 5, h * 5)
	// img.Fill(color.RGBA{225,255,255,255})

	return &Camera {
		img: img,
		Op: op,
		w: w,
		h: h,
		multiplier: 5,
	}
}

func (self *Camera) Translate(x, y float64) {
	self.Op.GeoM.Translate(x, y)
	self.Offset.X += x
	self.Offset.Y += y
}

func (self *Camera) Rotate(theta float64) {
	self.Op.GeoM.Rotate(theta)
}

func (self *Camera) Reset() {
	self.Op = &ebiten.DrawImageOptions{}
	self.Op.GeoM.Translate(-float64(self.w * self.multiplier / 3), -float64(self.h * self.multiplier / 3))
	self.Offset = FloatPoint{}
}

func(self *Camera) Draw(canvas *ebiten.Image) {
	canvas.DrawImage(self.img, self.Op)
	// I'm doing this to clear the screen between each draw
	// so the last frame doesn't appear on it
	self.img.Clear()
}

func (self *Camera) DrawImage(scene *ebiten.Image, op ebiten.DrawImageOptions) {
	op.GeoM.Translate(float64(self.w * self.multiplier / 3), float64(self.h * self.multiplier / 3))
	self.img.DrawImage(scene, &op)
}

func (self *Camera) Set(x, y int, color color.Color) {
	self.img.Set((self.w * self.multiplier / 3) + x, (self.h * self.multiplier / 3) + y, color)
}
