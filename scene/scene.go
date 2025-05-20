package scene

import (
	"github.com/eagledb14/guardsman/actor"
	"github.com/eagledb14/guardsman/camera"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type IScene interface {
	Draw(*ebiten.Image)
	Update() error
}


type DungeonScene struct {
	cam *camera.Camera
	State *GameState
	board *Board
	w, h int
	player actor.Actor
}

func NewDungeonScene(w, h int, state *GameState) *DungeonScene {
	cam := camera.NewCamera(w, h)
	d := DungeonScene {
		cam: cam,
		State: state,
		w: w,
		h: h,
		player: actor.NewPlayerActor(0, 0, 0),
	}

	d.nextBoard()
	return &d
}

func (self *DungeonScene) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		self.nextBoard()
	}
	action := self.player.Update()
	_ = action

	self.cam.Translate(-action.Translate.X, -action.Translate.Y)

	return nil
}

func (self *DungeonScene) nextBoard() {
	self.board = NewBoard(30,30,100)
	self.cam.Reset()
	x, y := self.board.Offset()
	self.cam.Translate(float64(self.w / 2) - y, float64(self.h / 2) - x)

	self.player.Op.GeoM.Reset()
	self.player.Op.GeoM.Translate(y, x)
}

func (self *DungeonScene) Draw(canvas *ebiten.Image) {
	self.board.Draw(self.cam)
	self.player.Draw(self.cam)
	self.cam.Draw(canvas)
}
