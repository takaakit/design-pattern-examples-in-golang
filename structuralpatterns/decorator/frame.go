// ˅
package main

// ˄

type Frame struct {
	// ˅

	// ˄

	Display

	display IDisplay

	// ˅

	// ˄
}

func NewFrame(display IDisplay) *Frame {
	// ˅
	frame := &Frame{}
	frame.display = display
	return frame
	// ˄
}

// ˅

// ˄
