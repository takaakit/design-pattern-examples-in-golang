package flyweight

import (
	"fmt"
	"os"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"

	"testing"
)

/*
Display a string consisting of large characters (0-9 digits only).
Large character objects are not created until they are needed.
And the created objects are reused.

Example Output
-----
Please enter digits (ex. 1212123): 123

     ####
      ###
      ###
      ###
      ###
      ###
    #######



   ########
         ###
         ###
   ########
  #
  #
  ##########



   ########
         ###
         ###
   ########
         ###
  #      ###
   ########
*/

func TestMain(m *testing.M) {
	var edit *walk.LineEdit
	var pb *walk.PushButton

	_, err := declarative.MainWindow{
		Title:   "Input Dialog",
		MinSize: declarative.Size{Width: 200, Height: 50},
		Size:    declarative.Size{Width: 200, Height: 50},
		Layout:  declarative.VBox{},
		Children: []declarative.Widget{
			declarative.Label{
				Text: "Please enter digits (ex. 1212123):",
			},
			declarative.LineEdit{
				AssignTo: &edit,
				MinSize:  declarative.Size{Width: 110, Height: 0},
			},
			declarative.PushButton{
				AssignTo: &pb,
				Text:     "OK",
				OnClicked: func() {
					lss := NewLargeSizeString(edit.Text())
					lss.Display()

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
