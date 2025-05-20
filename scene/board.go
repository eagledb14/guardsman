package scene

import (
	"image/color"

	"github.com/eagledb14/guardsman/graphics"
	"github.com/eagledb14/guardsman/camera"

	"github.com/hajimehoshi/ebiten/v2"
)


type Board struct {
	tiles [][]int
	w, h int
	length int
	start Point
	end Point
	size int
}

type Point struct {
	X,Y int
}

func NewBoard(w, h, length int) *Board {
	tiles := make([][]int, h)
	for i := range tiles {
		tiles[i] = make([]int, w)
	}

	b := &Board{
		w: w,
		h: h,
		length: length,
		tiles: tiles,
		size: 80,
	}

	Walk(&b.tiles, length)
	start := findStart(&b.tiles)
	end := findEnd(&b.tiles, start)
	b.start = start
	b.end = end


	b.tiles[b.start.Y][b.start.X] = 2
	b.tiles[b.end.Y][b.end.X] = 3

	return b
}

func (self *Board) Draw(camera *camera.Camera) {
	for i := range self.w {
		for j := range self.h {

			img := ebiten.NewImage(1,1)

			//floor
			if self.tiles[j][i] == 1 {
				img = graphics.NewRect(self.size, self.size, color.RGBA{155,176,195,255})
			// } else if self.tiles[j][i] == 0 { //walls
			// 	img = graphics.NewRect(self.size, self.size, color.RGBA{0,0,0,0})
			} else if self.tiles[j][i] == 2 { // start
				// img = graphics.NewRect(self.size, self.size, color.RGBA{0,255,0,255})
				img = graphics.NewRect(self.size, self.size, color.RGBA{155,176,195,255})
			} else if self.tiles[j][i] == 3 { //end
				img = graphics.NewRect(self.size, self.size, color.RGBA{255,0,0,255})
			}
			

			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((j + 1) * self.size), float64((i + 1) * self.size))

			camera.DrawImage(img, op)
		}
	}

	// draw border
	img := graphics.NewRect(self.size, self.size, color.RGBA{255,255,255,255})
	for i := range self.w + 1 {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(i * self.size), 0)
		camera.DrawImage(img, op)

		op = ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(i * self.size), float64((self.h + 1) * self.size))
		camera.DrawImage(img, op)
	}

	for i := range self.h + 1 {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, float64(i * self.size))
		camera.DrawImage(img, op)

		op = ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((self.w + 1) * self.size), float64(i * self.size))
		camera.DrawImage(img, op)
	}
}


func (self *Board) Offset() (float64, float64) {
	return float64(self.size * self.start.X) + float64(self.size), float64(self.size * self.start.Y) + float64(self.size)
}
