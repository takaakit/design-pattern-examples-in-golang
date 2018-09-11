// ˅
package main

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

func NewSideFrame(display IDisplay, frameChar string) *SideFrame {
	// ˅
	sideFrame := &SideFrame{}
	sideFrame.frameChar = frameChar
	sideFrame.Frame = *NewFrame(display)
	return sideFrame
	// ˄
}

func (self *SideFrame) GetLineText(row int) string {
	// ˅
	return self.frameChar + self.display.GetLineText(row) + self.frameChar
	// ˄
}

func (self *SideFrame) GetColumns() int {
	// ˅
	return 1 + self.display.GetColumns() + 1
	// ˄
}

func (self *SideFrame) GetRows() int {
	// ˅
	return self.display.GetRows()
	// ˄
}

// ˅

// ˄
