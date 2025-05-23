package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Actor struct {
	img *ebiten.Image
	Pos FloatPoint
	Rotation float64
	Controller IController
	// move bool
	size int
	Weapon IWeopan
}

type FloatPoint struct {
	X, Y float64
}

func NewActor(x, y float64, rotation float64, controller IController) *Actor {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.GeoM.Rotate(rotation)
	size := 50

	a := &Actor {
		img: NewRect(size, size, color.RGBA{0,102,204,255}),
		Controller: controller,
		size: size,
	}
	a.Weapon = NewPistol(a)
	return a
}

func NewPlayerActor(x, y float64, rotation float64) *Actor {
	c := NewPlayerController()

	return NewActor(x, y, rotation, c)
}

func (self *Actor) Update() Action {
	a := self.Controller.Update(self.Pos)

	oldPos := self.Pos
	self.Pos.X += a.Translate.X
	self.Pos.Y += a.Translate.Y

	// if self.Controller.
	if hitWall(self.GetHitBox()) {
		self.Pos = oldPos
		a.Translate = FloatPoint{}
	}

	self.Rotation = a.Rotate

	if a.DoShoot {
		self.Weapon.Fire()
	}
	self.Weapon.Update()

	return a
}

func (self *Actor) Reset() {
	self.Pos = FloatPoint{}
}

func (self *Actor) Translate(x, y float64) {
	self.Pos.X += x
	self.Pos.Y += y
}

func (self *Actor) Draw(cam *Camera) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(self.size / 2), -float64(self.size / 2))
	op.GeoM.Rotate(self.Rotation)
	op.GeoM.Translate(self.Pos.X , self.Pos.Y)

	self.Weapon.Draw(cam)

	cam.DrawImage(self.img, op)
}

func (self *Actor) GetHitBox() Hitbox {
	return NewHitBox(
		self.Pos.X, 
		self.Pos.Y,
		float64(self.size),
		float64(self.size),
	)
}
