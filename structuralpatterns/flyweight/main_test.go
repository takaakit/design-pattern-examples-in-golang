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

func Test(t *testing.T) {
	var le *walk.LineEdit
	var pb *walk.PushButton
	var mw *walk.MainWindow

	_, err := declarative.MainWindow{
		AssignTo: &mw,
		Title:    "Input Dialog",
		MinSize:  declarative.Size{Width: 200, Height: 50},
		Size:     declarative.Size{Width: 200, Height: 50},
		Layout:   declarative.VBox{},
		Children: []declarative.Widget{
			declarative.Label{
				Text: "Please enter digits (ex. 1212123):",
			},
			declarative.LineEdit{
				AssignTo: &le,
				MinSize:  declarative.Size{Width: 110, Height: 0},
			},
			declarative.PushButton{
				AssignTo: &pb,
				Text:     "OK",
				OnClicked: func() {
					lss := NewLargeSizeString(le.Text())
					lss.Display()

					mw.Close()
				},
			},
		},
	}.Run()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
