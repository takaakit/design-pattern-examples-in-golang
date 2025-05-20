package builder

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"

	"testing"
)

/*
Create documents in HTML format and text format. It is possible to create different documents
in the same construction process.
*/

func Test(t *testing.T) {
	var pb1 *walk.PushButton
	var pb2 *walk.PushButton

	_, err := declarative.MainWindow{
		Title:   "Choose 1 or 2",
		MinSize: declarative.Size{Width: 200, Height: 50},
		Size:    declarative.Size{Width: 200, Height: 50},
		Layout:  declarative.VBox{},
		Children: []declarative.Widget{
			declarative.Label{
				Text: "Choose 1 or 2:",
			},
			declarative.PushButton{
				AssignTo: &pb1,
				Text:     "1: Create a plain text document",
				OnClicked: func() {
					plainTextBuilder := NewPlainTextBuilder()
					director := NewDirector(plainTextBuilder)
					director.Build()
					content := plainTextBuilder.GetContent()
					fmt.Println(content)

					os.Exit(0)
				},
			},
			declarative.PushButton{
				AssignTo: &pb2,
				Text:     "2: Create an html file document",
				OnClicked: func() {
					htmlBuilder := NewHTMLBuilder()
					director := NewDirector(htmlBuilder)
					director.Build()
					fileName := htmlBuilder.GetFileName()
					fmt.Println(fileName + " has been created.")
					currentDir, _ := os.Getwd()
					fmt.Println("Output File: " + filepath.Join(currentDir, fileName))

					os.Exit(0)
				},
			},
		},
	}.Run()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
