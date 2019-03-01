// ˅
package main

import (
	"fmt"
	"os"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// ˄

type AppMain struct {
	// ˅

	// ˄

	buttonUndo *walk.PushButton

	buttonCancel *walk.PushButton

	// Painting history
	history *HistoryCommand

	canvas *PaintingCanvas

	// ˅

	// ˄
}

func NewAppMain() *AppMain {
	// ˅
	var mw *walk.MainWindow
	var pb1 *walk.PushButton
	var pb2 *walk.PushButton

	appMain := &AppMain{}
	appMain.history = NewHistoryCommand()
	appMain.canvas = NewPaintingCanvas(mw)
	appMain.buttonUndo = pb1
	appMain.buttonCancel = pb2

	mainWindow := MainWindow{
		AssignTo: &appMain.canvas.window,
		Title:    "Command Example",
		Size:     Size{480, 360},
		OnMouseMove: func(x int, y int, btn walk.MouseButton) {
			if btn == walk.LeftButton {
				appMain.onDragged(x, y)
			}
		},
		Layout: HBox{Margins: Margins{0, 310, 0, 0}},
		Children: []Widget{
			PushButton{
				AssignTo:  &appMain.buttonUndo,
				Text:      "Undo",
				OnClicked: appMain.undo,
			},
			PushButton{
				AssignTo:  &appMain.buttonCancel,
				Text:      "Cancel",
				OnClicked: appMain.cancel,
			},
		},
	}

	if _, err := mainWindow.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return appMain
	// ˄
}

func (self *AppMain) onDragged(paintingPosX int, paintingPosY int) {
	// ˅
	paintingCommand := NewPaintingCommand(self.canvas, paintingPosX, paintingPosY)
	self.history.Add(paintingCommand)
	paintingCommand.Execute()
	// ˄
}

func (self *AppMain) cancel() {
	// ˅
	self.canvas.Clear()
	self.history.Clear()
	// ˄
}

func (self *AppMain) undo() {
	// ˅
	self.canvas.Clear()
	self.history.Undo()
	self.history.Execute()
	// ˄
}

// ˅

// ˄
