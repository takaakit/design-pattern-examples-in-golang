// ˅
package main

import (
	"os"

	"github.com/lxn/walk"
)

// ˄

type PaintingCanvas struct {
	// ˅

	// ˄

	paintingColor walk.Color

	// Radius of the painting point
	pointRadius int

	window *walk.MainWindow

	// ˅

	// ˄
}

func NewPaintingCanvas(window *walk.MainWindow) *PaintingCanvas {
	// ˅
	paintingCanvas := &PaintingCanvas{}
	paintingCanvas.window = window
	paintingCanvas.pointRadius = 6.0
	paintingCanvas.paintingColor = walk.RGB(0, 0, 255)

	return paintingCanvas
	// ˄
}

func (self *PaintingCanvas) Paint(paintingPosX int, paintingPosY int) {
	// ˅
	canvas, _ := self.window.CreateCanvas()

	linesBrush, err := walk.NewSolidColorBrush(self.paintingColor)
	if err != nil {
		os.Exit(1)
	}
	defer linesBrush.Dispose()

	linesPen, err := walk.NewGeometricPen(walk.PenDash, self.pointRadius*2, linesBrush)
	if err != nil {
		os.Exit(1)
	}
	defer linesPen.Dispose()

	if err := canvas.DrawLine(linesPen, walk.Point{paintingPosX, paintingPosY}, walk.Point{paintingPosX, paintingPosY}); err != nil {
		os.Exit(1)
	}
	// ˄
}

func (self *PaintingCanvas) Clear() {
	// ˅
	canvas, _ := self.window.CreateCanvas()

	bounds := self.window.ClientBounds()

	ellipseBrush, err := walk.NewSolidColorBrush(walk.RGB(255, 255, 255))
	if err != nil {
		os.Exit(1)
	}
	defer ellipseBrush.Dispose()

	if err := canvas.FillRectangle(ellipseBrush, bounds); err != nil {
		os.Exit(1)
	}
	// ˄
}

// ˅

// ˄
