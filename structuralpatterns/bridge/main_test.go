package bridge

import (
	"testing"
)

/*
Display only one line or display the specified number of lines.
*/

func TestMain(m *testing.M) {
	d1 := NewDisplay(NewTextDisplayImpl("Japan"))
	d1.Output()

	d2 := NewMultiLineDisplay(NewTextDisplayImpl("The United States of America"))
	d2.Output()
	d2.OutputMultiple(3)
}
