package main

import (

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type IScene interface {
	Draw(*ebiten.Image)
	Update() error
}


type DungeonScene struct {
	State *GameState
}

func NewDungeonScene() *DungeonScene {
	d := DungeonScene {}

	state.NextBoard()
	return &d
}

func (self *DungeonScene) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.NextBoard()
	}
	action := state.player.Update()

	state.Cam.Translate(-action.Translate.X, -action.Translate.Y)

	for _, bullet := range state.Bullets {
		bullet.Update()
	}

	return nil
}

func (self *DungeonScene) Draw(canvas *ebiten.Image) {
	state.board.Draw(state.Cam)

	for _, bullet := range state.Bullets {
		bullet.Draw(state.Cam)
	}

	state.player.Draw(state.Cam)

	state.Cam.Draw(canvas)
}
