package main

import (
	"image/color"
)

type Hitbox struct {
	X float64
	Y float64
	Width  float64
	Height float64
}

func NewHitBox(x, y, w, h float64) Hitbox {
	offset := state.Cam.Offset
	return Hitbox{
		X: x + offset.X,
		Y: y + offset.Y,
		Width: w,
		Height: h,
	}
}

func (h1 *Hitbox) IsOverlap(h2 *Hitbox) bool {
	h1CenterX, h1CenterY := h1.GetCenter()
	xContains := h1CenterX >= h2.X && h1CenterX < h2.X+h2.Width
	yContains := h1CenterY >= h2.Y && h1CenterY < h2.Y+h2.Height

	return xContains && yContains
}

func (h1 *Hitbox) IsHit(h2 *Hitbox) bool {
	xOverlap := h1.X < h2.X+h2.Width && h1.X+h1.Width > h2.X
	yOverlap := h1.Y < h2.Y+h2.Height && h1.Y+h1.Height > h2.Y

	return xOverlap && yOverlap
}

func (h1 *Hitbox) Draw(cam *Camera) {
	h := int(h1.Width)
	w := int(h1.Height)
	x := int(h1.X)
	y := int(h1.Y)

	color := color.RGBA{255,0,0,255}

	// top and bottom
	for i := 0; i < w; i++ {
		cam.Set(0 + x, i + y, color)
		cam.Set(h - 1, i + y, color)
	}

	// left and right
	for i := 0; i < h; i++ {
		cam.Set(i + x, 0 + y, color)
		cam.Set(i + x, w - 0 + y, color)
	}
}

func (h *Hitbox) GetCenter() (centerX, centerY float64) {
	return h.X, h.Y
}
