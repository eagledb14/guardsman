package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)


type Board struct {
	tiles [][]int
	mesh *ebiten.Image
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
	b.mesh = ebiten.NewImage((w + 2)* b.size, (h + 2) * b.size)
	b.CreateMesh()

	return b
}

func (self *Board) Draw(camera *Camera) {
	op := ebiten.DrawImageOptions{}
	camera.DrawImage(self.mesh, op)
}

// func (self *Board) Draw(camera *camera.Camera) {
func (self *Board) CreateMesh() {
	for i := range self.w {
		for j := range self.h {

			img := ebiten.NewImage(1,1)

			//floor
			if self.tiles[j][i] == 1 {
				img = NewRect(self.size, self.size, color.RGBA{155,176,195,255})
			// } else if self.tiles[j][i] == 0 { //walls
			// 	img = graphics.NewRect(self.size, self.size, color.RGBA{0,0,0,0})
			} else if self.tiles[j][i] == 2 { // start
				// img = graphics.NewRect(self.size, self.size, color.RGBA{0,255,0,255})
				img = NewRect(self.size, self.size, color.RGBA{155,176,195,255})
			} else if self.tiles[j][i] == 3 { //end
				img = NewRect(self.size, self.size, color.RGBA{255,0,0,255})
			}
			
			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((j + 1) * self.size), float64((i + 1) * self.size))

			self.mesh.DrawImage(img, &op)
		}
	}

	// draw border
	// img := NewRect(self.size, self.size, color.RGBA{255,255,255,255})
	// for i := range self.w + 1 {
	// 	op := ebiten.DrawImageOptions{}
	// 	op.GeoM.Translate(float64(i * self.size), 0)
	// 	self.mesh.DrawImage(img, &op)
	//
	// 	op = ebiten.DrawImageOptions{}
	// 	op.GeoM.Translate(float64(i * self.size), float64((self.h + 1) * self.size))
	// 	self.mesh.DrawImage(img, &op)
	// }
	//
	// for i := range self.h + 1 {
	// 	op := ebiten.DrawImageOptions{}
	// 	op.GeoM.Translate(0, float64(i * self.size))
	// 	self.mesh.DrawImage(img, &op)
	//
	// 	op = ebiten.DrawImageOptions{}
	// 	op.GeoM.Translate(float64((self.w + 1) * self.size), float64(i * self.size))
	// 	self.mesh.DrawImage(img, &op)
	// }
}

func (self *Board) Offset() (float64, float64) {
	return float64(self.size * self.start.X) + float64(self.size), float64(self.size * self.start.Y) + float64(self.size)
}

func (self *Board) GetWallHitBox() []Hitbox {
	boxes := []Hitbox{}
	for i := range self.w {
		for j := range self.h {
			if self.tiles[j][i] == 0 {
				newBox := NewHitBox(float64((j + 1) * self.size),float64((i + 1) * self.size), float64(self.size), float64(self.size))
				boxes = append(boxes, newBox)
			}
		}
	}

	for i := range self.w + 1 {
		newBox := NewHitBox(float64(i * self.size), 0, float64(self.size), float64(self.size))
		boxes = append(boxes, newBox)

		newBox = NewHitBox(float64(i * self.size), float64(self.h + 1), float64(self.size), float64(self.size))
		boxes = append(boxes, newBox)
	}

	for i := range self.h + 1 {
		newBox := NewHitBox(0, float64(i * self.size), float64(self.size), float64(self.size))
		boxes = append(boxes, newBox)

		newBox = NewHitBox(float64(self.w + 1), float64(i * self.size), float64(self.size), float64(self.size))
		boxes = append(boxes, newBox)
	}

	return boxes
}

func (self *Board) IsOnFinish(playerBox Hitbox) bool {
	finishBox := NewHitBox(float64(self.end.Y * self.size + self.size), float64(self.end.X * self.size + self.size), float64(self.size), float64(self.size))

	return finishBox.IsOverlap(&playerBox)
}
