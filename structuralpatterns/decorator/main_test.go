package decorator

import (
	"testing"
)

/*
Display a string with decorative frames. The frames can be combined arbitrarily.
*/

func Test(t *testing.T) {
	displayA := NewMessageDisplay("Nice to meet you.")
	displayA.Show()

	displayB := NewSideFrame(NewMessageDisplay("Nice to meet you."), "!")
	displayB.Show()

	displayC := NewFullFrame(NewSideFrame(NewMessageDisplay("Nice to meet you."), "!"))
	displayC.Show()
}
