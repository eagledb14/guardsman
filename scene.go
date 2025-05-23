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
	_ = action

	state.Cam.Translate(-action.Translate.X, -action.Translate.Y)

	state.UpdateBullets()

	if state.board.IsOnFinish(state.player.GetCenterBox()) {
		state.NextBoard()
	}

	return nil
}

func (self *DungeonScene) Draw(canvas *ebiten.Image) {
	state.board.Draw(state.Cam)

	for _, bullet := range state.Bullets {
		bullet.Draw(state.Cam)
	}

	// for _, box := range state.board.GetWallHitBox() {
	// 	box.Draw(state.Cam)
	// }

	state.player.Draw(state.Cam)

	state.Cam.Draw(canvas)
}
