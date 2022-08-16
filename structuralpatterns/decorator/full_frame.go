// ˅
package decorator

import (
	"fmt"
	"strings"
)

// ˄

type FullFrame struct {
	// ˅

	// ˄

	Frame

	// ˅

	// ˄
}

func NewFullFrame(display Display) *FullFrame {
	// ˅
	return &FullFrame{Frame: *NewFrame(display)}
	// ˄
}

func (f *FullFrame) GetLineText(row int) string {
	// ˅
	if row == 0 {
		// Upper frame
		return "+" + f.createLine("-", f.display.GetColumns()) + "+"
	} else if row == f.display.GetRows()+1 {
		// Bottom frame
		return "+" + f.createLine("-", f.display.GetColumns()) + "+"
	} else {
		// Other
		return "|" + f.display.GetLineText(row-1) + "|"
	}
	// ˄
}

func (f *FullFrame) GetColumns() int {
	// ˅
	return 1 + f.display.GetColumns() + 1
	// ˄
}

func (f *FullFrame) GetRows() int {
	// ˅
	return 1 + f.display.GetRows() + 1
	// ˄
}

// Show all
func (f *FullFrame) Show() {
	// ˅
	for i := 0; i < f.GetRows(); i++ {
		fmt.Println(f.GetLineText(i))
	}
	// ˄
}

func (f *FullFrame) createLine(ch string, size int) string {
	// ˅
	return strings.Repeat(ch, size)
	// ˄
}

// ˅

// ˄
