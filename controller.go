package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type IController interface {
	Update(FloatPoint) Action
}


type Action struct {
	Translate FloatPoint
	Rotate float64
	DoShoot bool
}

type PlayerController struct {
	r float64
}

func NewPlayerController() *PlayerController {
	return &PlayerController {
	}
}


func (self *PlayerController) Update(pos FloatPoint) Action {
	a := Action{}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		a.Translate.Y -= 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		a.Translate.Y += 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		a.Translate.X -= 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		a.Translate.X += 5
	}

	x, y := ebiten.CursorPosition()
	dx := state.Cam.Offset.X
	dy := state.Cam.Offset.Y

	a.Rotate += math.Atan2(pos.Y - float64(y) + dy, pos.X - float64(x) + dx)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		a.DoShoot = true
	}

	return a
}

func hitWall(actorBox *resolv.ConvexPolygon) bool {
	walls := state.board.GetWallHitBox()

	for _, wall := range walls {
		if  actorBox.IsIntersecting(wall) {
			return true
		}
	}

	return false
}

type AIController struct {
}

func NewAIController() *PlayerController {
	return &PlayerController {
	}
}
