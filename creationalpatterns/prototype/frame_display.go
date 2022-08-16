// ˅
package prototype

import (
	"fmt"
	"strings"

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
	fmt.Println(strings.Repeat(f.borderChar, len(message)+4))
	fmt.Println(f.borderChar + " " + message + " " + f.borderChar)
	fmt.Println(strings.Repeat(f.borderChar, len(message)+4))
	// ˄
}

// ˅

// ˄
