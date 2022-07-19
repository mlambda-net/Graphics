package main

import (
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/sdlcanvas"
	"image"
)

type Window interface {
	Create(w, h int, name string) error
	Render(render func(graph Graph))
}

type window struct {
	render func(graph Graph)
	graph  *graph
	canvas *canvas.Canvas
	size   Size
	wnd    *sdlcanvas.Window
}

func (w *window) Render(render func(graph Graph)) {
	w.render = render
	defer w.wnd.Close()
	w.wnd.MainLoop(func() {
		width, height := float64(w.size.weigh), float64(w.size.height)
		w.canvas.SetFillStyle("#000")
		w.canvas.FillRect(0, 0, width, height)
		w.run()
	})
}

func (w *window) Create(weigh, height int, name string) error {
	wnd, cv, err := sdlcanvas.CreateWindow(weigh, height, name)
	if err != nil {
		return err
	}

	g := &graph{}

	wnd.MouseUp = func(button, x, y int) {
	}

	wnd.MouseMove = func(x, y int) {
		g.setPosition(Position{x: x, y: y})
	}

	wnd.MouseDown = func(button, x, y int) {
	}

	w.canvas = cv
	w.graph = g
	w.size = Size{weigh: weigh, height: weigh}
	w.wnd = wnd

	return nil
}

func (w *window) run() {
	if w.render != nil {
		w.graph.image = image.NewNRGBA(image.Rect(0, 0, w.size.weigh, w.size.height))
		w.render(w.graph)

		w.canvas.DrawImage(w.graph.image)
	}
}

func NewWindow() Window {
	return &window{}
}
