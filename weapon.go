package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)


type IWeopan interface {
	Fire()
	Update()
	Draw(*Camera)
}

type Pistol struct {
	actor *Actor
	cooldown int
}

func NewPistol(actor *Actor) *Pistol {
	return &Pistol {
		actor: actor,
	}
}

func (self *Pistol) Fire() {
	// create bullets and add to the gamestate struct
	if self.cooldown <= 0 {
		state.Bullets = append(state.Bullets, NewBullet(self.actor.Pos.X, self.actor.Pos.Y, self.actor.Rotation))
		// state.Bullets = append(state.Bullets, NewBullet(self.actor.Pos.X, self.actor.Pos.Y, self.actor.Rotation + 0.1))
		// state.Bullets = append(state.Bullets, NewBullet(self.actor.Pos.X, self.actor.Pos.Y, self.actor.Rotation - 0.1))
		self.cooldown = 50
	}
}

func (self * Pistol) Update() {
	self.cooldown -= 1
}

func (self *Pistol) Draw(cam *Camera) {

}


type Bullet struct {
	img *ebiten.Image
	pos FloatPoint
	vel FloatPoint
	size int
}

func NewBullet(x, y float64, rotation float64) *Bullet {
	size := 10

	dy := 10 * math.Sin(rotation)
	dx := 10 * math.Cos(rotation)

	return &Bullet{
		img: NewRect(size, size, color.RGBA{255, 234, 0, 255}),
		size: size,
		vel: FloatPoint{
			X: -dx,
			Y: -dy,
		},
		pos: FloatPoint{
			X: x,
			Y: y,
		},
	}
}

func (self *Bullet) Update() error {
	self.pos.X += self.vel.X
	self.pos.Y += self.vel.Y
	return nil
}

func (self *Bullet) Draw(cam *Camera) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(self.pos.X, self.pos.Y)

	cam.DrawImage(self.img, op)
}
