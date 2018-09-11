// ˅
package main

// ˄

type Display struct {
	// ˅

	// ˄

	impl DisplayImpl

	// ˅

	// ˄
}

func NewDisplay(impl DisplayImpl) *Display {
	// ˅
	return &Display{impl}
	// ˄
}

func (self *Display) Output() {
	// ˅
	self.Open()
	self.Write()
	self.Close()
	// ˄
}

func (self *Display) Open() {
	// ˅
	self.impl.ImplOpen()
	// ˄
}

func (self *Display) Write() {
	// ˅
	self.impl.ImplWrite()
	// ˄
}

func (self *Display) Close() {
	// ˅
	self.impl.ImplClose()
	// ˄
}

// ˅

// ˄
