package actor

import "github.com/hajimehoshi/ebiten/v2"

type IController interface {
	Update() Action
}


type Action struct {
	Translate Point
	Rotate float64
	DoShoot bool
}

type Point struct {
	X, Y float64
}


type PlayerController struct {
}

func NewPlayerController() *PlayerController {
	return &PlayerController {
	}
}

func (self *PlayerController) Update() Action {
	a := Action{}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		// self.cam.Translate(0, 10)
		a.Translate.Y -= 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		// self.cam.Translate(0, -10)
		a.Translate.Y += 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		// self.cam.Translate(10, 0)
		a.Translate.X -= 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		// self.cam.Translate(-10, 0)
		a.Translate.X += 5
	}

	return a
}
