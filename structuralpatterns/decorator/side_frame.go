// ˅
package decorator

import (
	"fmt"
	"os"
)

// ˄

type SideFrame struct {
	// ˅

	// ˄

	Frame

	// Decoration character
	frameChar string

	// ˅

	// ˄
}

func NewSideFrame(display Display, frameChar string) *SideFrame {
	// ˅
	if len(frameChar) != 1 {
		fmt.Println("Only one character is allowed in a side frame.")
		os.Exit(1)
	}
	return &SideFrame{Frame: *NewFrame(display), frameChar: frameChar}
	// ˄
}

func (s *SideFrame) GetLineText(row int) string {
	// ˅
	return s.frameChar + s.display.GetLineText(row) + s.frameChar
	// ˄
}

func (s *SideFrame) GetColumns() int {
	// ˅
	return 1 + s.display.GetColumns() + 1
	// ˄
}

func (s *SideFrame) GetRows() int {
	// ˅
	return s.display.GetRows()
	// ˄
}

// Show all
func (s *SideFrame) Show() {
	// ˅
	for i := 0; i < s.GetRows(); i++ {
		fmt.Println(s.GetLineText(i))
	}
	// ˄
}

// ˅

// ˄
