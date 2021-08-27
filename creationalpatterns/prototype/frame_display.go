// ˅
package prototype

import (
	"fmt"
	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/prototype/framework"
)

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
	return &FrameDisplay{borderChar: borderChar}
	// ˄
}

func (f *FrameDisplay) Clone() Display {
	// ˅
	return NewFrameDisplay(f.borderChar)
	// ˄
}

func (f *FrameDisplay) Show(message string) {
	// ˅
	var length = len(message)
	for i := 0; i < length+4; i++ {
		fmt.Print(f.borderChar)
	}
	fmt.Println("")
	fmt.Println(f.borderChar + " " + message + " " + f.borderChar)
	for i := 0; i < length+4; i++ {
		fmt.Print(f.borderChar)
	}
	fmt.Println("")
	// ˄
}

// ˅

// ˄
