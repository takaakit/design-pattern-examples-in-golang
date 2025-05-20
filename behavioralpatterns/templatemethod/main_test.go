package templatemethod

import "testing"

/*
Display a character or string repeatedly 5 times.
*/

func Test(t *testing.T) {
	display1 := NewCharDisplay("H")
	display1.Output(display1)

	display2 := NewStringDisplay("Hello world.")
	display2.Output(display2)
}
