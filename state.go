package main

import (
)

var state GameState

type GameState struct {
	Cam *Camera
	player Actor
	board *Board
	w, h int
}

// func NewGameState() GameState {
func InitGameState(w, h int) {
	cam := NewCamera(w, h)
	state = GameState{
		Cam: cam,
		player: NewPlayerActor(0, 0, 0),
		w: w,
		h: h,
	}
}

func (self *GameState) NextBoard() {
	self.board = NewBoard(30,30,100)
	state.Cam.Reset()
	x, y := self.board.Offset()
	state.Cam.Translate(float64(self.w / 2) - y, float64(self.h / 2) - x)
	// self.cam.Translate(0,0)

	state.player.Reset()
	state.player.Translate(y,x)
	// self.player.Translate(0,0)
}
