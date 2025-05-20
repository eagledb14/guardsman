package main

import (
	// "image/color"

	// "github.com/eagledb14/guardsman/camera"
	// "github.com/eagledb14/guardsman/graphics"
	"github.com/eagledb14/guardsman/scene"
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/tinne26/etxt"
	// "github.com/tinne26/fonts/liberation/lbrtmono"
)

type Game struct {
	// text *etxt.Renderer
	scene scene.IScene
}

func (self *Game) Layout(w, h int) (int, int) {
	return 1920, 1080
}

func (self *Game) Update() error {
	return self.scene.Update()
	// img := graphics.NewRect(100, 100, color.RGBA{255, 255, 0, 255})
	// self.cam.DrawImage(img, op)

	// return nil
}

func (self *Game) Draw(canvas *ebiten.Image) {
	self.scene.Draw(canvas)
	// self.cam.Draw(canvas)
}

func main() {
	g := &Game{}
	s := &scene.GameState{}
	_ = s
	// c := camera.NewCamera(640, 480)
	// for i := range 100 {
	// 	for j := range 100 {
	// 		op := ebiten.DrawImageOptions{}
	// 		op.GeoM.Translate(float64(i * 150), float64(j * 150))
	// 		img := graphics.NewRect(100, 100, color.RGBA{uint8(i * 2), uint8(j * 2), 0, 255})
	// 		c.DrawImage(img, op)
	// 	}
	// }
	//
	// g.cam = c
	x,y := g.Layout(0,0)
	g.scene = scene.NewDungeonScene(x, y, s)
	ebiten.SetWindowSize(640,480)
	// ebiten.SetWindowSize(1000, 1000)
	// b := scene.NewBoard(30, 30, 10)
	// g := &Game{}
	// g.board = b
	// g.cam = ebiten.NewImage(680, 480)
	// g.camOp = &ebiten.DrawImageOptions{}


	// b.GenerateDungeon()
	// b.DrawDungeon()
	// fmt.Println(b)
	// renderer := etxt.NewRenderer()
	// renderer.SetFont(lbrtmono.Font())
	// renderer.Utils().SetCache8MiB()
	//
	// renderer.SetColor(color.RGBA{239, 91, 91, 255})
	// renderer.SetAlign(etxt.Center)
	// renderer.SetSize(72)
	ebiten.RunGame(g)
}
