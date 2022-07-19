package main

import (
	"image"
	"image/color"
)

type Graph interface {
	Position() (x, y int)
	DrawRectangle(x, y, w, h int, color color.Color)
}

type graph struct {
	position Position
	image    *image.NRGBA
}

func (g *graph) DrawRectangle(x, y, w, h int, color color.Color) {
	for i := 1; i < w; i++ {
		for j := 1; j < h; j++ {
			g.image.Set(x+i, y+j, color)
		}
	}
}

func (g *graph) setPosition(p Position) {
	g.position = p
}

func (g *graph) Position() (x, y int) {
	return g.position.x, g.position.y
}
