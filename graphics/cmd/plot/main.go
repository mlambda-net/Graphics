package main

import (
	"image/color"
)

func main() {

	window := NewWindow()
	err := window.Create(500, 500, "My Paint")
	if err != nil {
		panic(err)
	}

	window.Render(func(graph Graph) {
		x, y := graph.Position()
		graph.DrawRectangle(x, y, 100, 100, color.White)
	})

}
