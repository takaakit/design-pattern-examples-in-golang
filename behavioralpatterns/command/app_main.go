// ˅
package command

import (
	"fmt"
	"os"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

// ˄

type AppMain struct {
	// ˅

	// ˄

	// Painting history
	history *HistoryCommand

	canvas *PaintingCanvas

	buttonUndo *walk.PushButton

	buttonClear *walk.PushButton

	// ˅

	// ˄
}

func NewAppMain() *AppMain {
	// ˅
	var mw *walk.MainWindow
	var pb1 *walk.PushButton
	var pb2 *walk.PushButton

	appMain := &AppMain{
		buttonUndo:  pb1,
		buttonClear: pb2,
		history:     NewHistoryCommand(),
		canvas:      NewPaintingCanvas(mw),
	}

	if _, err := (declarative.MainWindow{
		AssignTo: &appMain.canvas.window,
		Title:    "Command Example",
		MinSize:  declarative.Size{Width: 400, Height: 300},
		Size:     declarative.Size{Width: 400, Height: 300},
		OnMouseMove: func(x int, y int, btn walk.MouseButton) {
			if btn == walk.LeftButton {
				appMain.onDragged(x, y)
			}
		},
		Layout: declarative.HBox{Margins: declarative.Margins{Left: 0, Top: 300, Right: 0, Bottom: 0}},
		Children: []declarative.Widget{
			declarative.PushButton{
				AssignTo:  &appMain.buttonUndo,
				Text:      "Undo",
				OnClicked: appMain.undo,
			},
			declarative.PushButton{
				AssignTo:  &appMain.buttonClear,
				Text:      "Clear",
				OnClicked: appMain.clear,
			},
		},
	}.Run()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return appMain
	// ˄
}

func (a *AppMain) onDragged(x int, y int) {
	// ˅
	paintingCommand := NewPaintingCommand(a.canvas, x, y)
	a.history.Add(paintingCommand)
	paintingCommand.Execute()
	// ˄
}

func (a *AppMain) undo() {
	// ˅
	a.canvas.Clear()
	a.history.Undo()
	a.history.Execute()
	// ˄
}

func (a *AppMain) clear() {
	// ˅
	a.canvas.Clear()
	a.history.Clear()
	// ˄
}

// ˅

// ˄
