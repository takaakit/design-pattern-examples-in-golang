// ˅
package main

// ˄

type MultiLineDisplay struct {
	// ˅

	// ˄

	Display

	// ˅

	// ˄
}

func NewMultiLineDisplay(impl DisplayImpl) *MultiLineDisplay {
	// ˅
	multiLineDisplay := &MultiLineDisplay{}
	multiLineDisplay.Display = *NewDisplay(impl)
	return multiLineDisplay
	// ˄
}

// Repeat display for the specified number of times
func (self *MultiLineDisplay) OutputMultiple(times int) {
	// ˅
	self.Open()
	for i := 0; i < times; i++ {
		self.Write()
	}
	self.Close()
	// ˄
}

// ˅

// ˄
