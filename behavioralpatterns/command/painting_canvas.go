// ˅
package command

import (
	"os"

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

	linesBrush, err := walk.NewSolidColorBrush(p.paintingColor)
	if err != nil {
		os.Exit(1)
	}
	defer linesBrush.Dispose()

	linesPen, err := walk.NewGeometricPen(walk.PenDash, p.pointRadius*2, linesBrush)
	if err != nil {
		os.Exit(1)
	}
	defer linesPen.Dispose()

	if err := canvas.DrawLinePixels(linesPen, walk.Point{X: x, Y: y}, walk.Point{X: x, Y: y}); err != nil {
		os.Exit(1)
	}
	// ˄
}

func (p *PaintingCanvas) Clear() {
	// ˅
	canvas, _ := p.window.CreateCanvas()

	bounds := p.window.ClientBoundsPixels()

	ellipseBrush, err := walk.NewSolidColorBrush(walk.RGB(255, 255, 255))
	if err != nil {
		os.Exit(1)
	}
	defer ellipseBrush.Dispose()

	if err := canvas.FillRectanglePixels(ellipseBrush, bounds); err != nil {
		os.Exit(1)
	}
	// ˄
}

// ˅

// ˄
