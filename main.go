package main

import (
	// "image/color"


	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/tinne26/etxt"
	// "github.com/tinne26/fonts/liberation/lbrtmono"
)


type Game struct {
	scene IScene
}

func (self *Game) Layout(w, h int) (int, int) {
	return 1920, 1080
}

func (self *Game) Update() error {
	return self.scene.Update()
}

func (self *Game) Draw(canvas *ebiten.Image) {
	self.scene.Draw(canvas)
}

func main() {
	g := &Game{}
	s := &GameState{}
	_ = s

	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	x,y := g.Layout(0,0)
	InitGameState(x,y)
	g.scene = NewDungeonScene()
	ebiten.SetWindowSize(640,480)

	ebiten.RunGame(g)
}
