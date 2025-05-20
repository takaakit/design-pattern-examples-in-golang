package prototype

import (
	"testing"

	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/prototype/framework"
)

/*
Display a string enclosed with a frame line, or drawn with an underline. The Client (Main)
registers instances of the Display subclass in the Manager class. When necessary,
the Manager class asks those registered instances to return a clone. The Client (Main)
requires the returned clones to display.
*/

func Test(t *testing.T) {
	manager := NewManager()

	// Register instances of the "Display" subclass
	emphasisUnderline := NewUnderlineDisplay("~")
	manager.RegisterDisplay("emphasis", emphasisUnderline)
	highlightFrame := NewFrameDisplay("+")
	manager.RegisterDisplay("highlight", highlightFrame)
	warningFrame := NewFrameDisplay("#")
	manager.RegisterDisplay("warning", warningFrame)

	// Require to display
	display1 := manager.GetDisplay("emphasis")
	display1.Show("Nice to meet you.")
	display2 := manager.GetDisplay("highlight")
	display2.Show("Nice to meet you.")
	display3 := manager.GetDisplay("warning")
	display3.Show("Nice to meet you.")
}
