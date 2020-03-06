// ˅
package main

import "fmt"

// ˄

type FrameDisplay struct {
	// ˅

	// ˄

	borderChar string

	// ˅

	// ˄
}

func NewFrameDisplay(borderChar string) *FrameDisplay {
	// ˅
	return &FrameDisplay{borderChar}
	// ˄
}

func (self *FrameDisplay) CreateClone() Display {
	// ˅
	return NewFrameDisplay(self.borderChar)
	// ˄
}

func (self *FrameDisplay) Show(message string) {
	// ˅
	var length = len(message)
	for i := 0; i < length+4; i++ {
		fmt.Print(self.borderChar)
	}
	fmt.Println("")
	fmt.Println(self.borderChar + " " + message + " " + self.borderChar)
	for i := 0; i < length+4; i++ {
		fmt.Print(self.borderChar)
	}
	fmt.Println("")
	// ˄
}

// ˅

// ˄
