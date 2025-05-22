package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Actor struct {
	img *ebiten.Image
	pos FloatPoint
	rotation float64
	Controller IController
	move bool
	size int
}

type FloatPoint struct {
	X, Y float64
}

func NewActor(x, y float64, rotation float64, controller IController) Actor {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.GeoM.Rotate(rotation)
	size := 50

	return Actor {
		img: NewRect(size, size, color.RGBA{0,102,204,255}),
		Controller: controller,
		size: size,
	}
}

func NewPlayerActor(x, y float64, rotation float64) Actor {
	c := NewPlayerController()

	return NewActor(x, y, rotation, c)
}

func (self *Actor) Update() Action {
	a := self.Controller.Update(self.pos)

	// if !self.move {
	self.pos.X += a.Translate.X
	self.pos.Y += a.Translate.Y
	// }

	self.rotation = a.Rotate

	return a
}

func (self *Actor) Reset() {
	self.pos = FloatPoint{}
}

func (self *Actor) Translate(x, y float64) {
	self.pos.X += x
	self.pos.Y += y
}

func (self *Actor) Draw(cam *Camera) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(self.size / 2), -float64(self.size / 2))
	op.GeoM.Rotate(self.rotation)
	op.GeoM.Translate(self.pos.X , self.pos.Y)

	cam.DrawImage(self.img, op)
}
