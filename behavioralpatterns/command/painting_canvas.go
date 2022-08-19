// ˅
package command

import (
	"github.com/lxn/walk"
)

// ˄

type PaintingCanvas struct {
	// ˅

	// ˄

	// Radius of the painting point
	pointRadius int

	window *walk.MainWindow

	paintingColor walk.Color

	// ˅

	// ˄
}

func NewPaintingCanvas(window *walk.MainWindow) *PaintingCanvas {
	// ˅
	return &PaintingCanvas{
		paintingColor: walk.RGB(128, 255, 128),
		pointRadius:   10.0,
		window:        window,
	}
	// ˄
}

func (p *PaintingCanvas) Paint(x int, y int) {
	// ˅
	canvas, _ := p.window.CreateCanvas()
	linesBrush, _ := walk.NewSolidColorBrush(p.paintingColor)
	linesPen, _ := walk.NewGeometricPen(walk.PenDash, p.pointRadius*2, linesBrush)
	_ = canvas.DrawLinePixels(linesPen, walk.Point{X: x, Y: y}, walk.Point{X: x, Y: y})
	// ˄
}

func (p *PaintingCanvas) Clear() {
	// ˅
	canvas, _ := p.window.CreateCanvas()
	bounds := p.window.ClientBoundsPixels()
	ellipseBrush, _ := walk.NewSolidColorBrush(walk.RGB(255, 255, 255))
	_ = canvas.FillRectanglePixels(ellipseBrush, bounds)
	// ˄
}

// ˅

// ˄
