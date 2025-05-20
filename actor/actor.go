package actor

import (
	"image/color"

	"github.com/eagledb14/guardsman/camera"
	"github.com/eagledb14/guardsman/graphics"
	"github.com/hajimehoshi/ebiten/v2"
)

type Actor struct {
	img *ebiten.Image
	Op *ebiten.DrawImageOptions
	Controller IController
	move bool
}

func NewActor(x, y float64, rotation float64, controller IController) Actor {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.GeoM.Rotate(rotation)

	return Actor {
		Op: op,
		img: graphics.NewRect(10, 10, color.RGBA{0,102,204,255}),
		Controller: controller,
	}
}

func NewPlayerActor(x, y float64, rotation float64) Actor {
	c := NewPlayerController()

	return NewActor(x, y, rotation, c)
}


func (self *Actor) Update() Action {
	a := self.Controller.Update()

	// if !self.move {
	self.Op.GeoM.Translate(a.Translate.X, a.Translate.Y)
	// }

	self.Op.GeoM.Rotate(a.Rotate)

	return a
}

func (self *Actor) Draw(cam *camera.Camera) {
	cam.DrawImage(self.img, *self.Op)
}
