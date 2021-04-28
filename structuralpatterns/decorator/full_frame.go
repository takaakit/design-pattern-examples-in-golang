// ˅
package main

// ˄

type FullFrame struct {
	// ˅

	// ˄

	Frame

	// ˅

	// ˄
}

func NewFullFrame(display IDisplay) *FullFrame {
	// ˅
	fullFrame := &FullFrame{}
	fullFrame.Frame = *NewFrame(display)
	return fullFrame
	// ˄
}

func (self *FullFrame) GetLineText(row int) string {
	// ˅
	if row == 0 {
		// Upper frame
		return "+" + self.createLine("-", self.display.GetColumns()) + "+"
	} else if row == self.display.GetRows()+1 {
		// Bottom frame
		return "+" + self.createLine("-", self.display.GetColumns()) + "+"
	} else {
		// Other
		return "|" + self.display.GetLineText(row-1) + "|"
	}
	// ˄
}

func (self *FullFrame) GetColumns() int {
	// ˅
	return 1 + self.display.GetColumns() + 1
	// ˄
}

func (self *FullFrame) GetRows() int {
	// ˅
	return 1 + self.display.GetRows() + 1
	// ˄
}

func (self *FullFrame) createLine(ch string, size int) string {
	// ˅
	var buf string = ""
	for i := 0; i < size; i++ {
		buf += ch
	}
	return buf
	// ˄
}

// ˅

// ˄
