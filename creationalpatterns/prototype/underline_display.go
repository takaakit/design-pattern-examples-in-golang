// ˅
package prototype

import (
	"fmt"
	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/prototype/framework"
)

// ˄

type UnderlineDisplay struct {
	// ˅

	// ˄

	underlineChar string

	// ˅

	// ˄
}

func NewUnderlineDisplay(underlineChar string) *UnderlineDisplay {
	// ˅
	return &UnderlineDisplay{underlineChar: underlineChar}
	// ˄
}

func (u *UnderlineDisplay) Clone() Display {
	// ˅
	return NewUnderlineDisplay(u.underlineChar)
	// ˄
}

func (u *UnderlineDisplay) Show(message string) {
	// ˅
	var length = len(message)
	fmt.Println("\"" + message + "\"")
	fmt.Print(" ")
	for i := 0; i < length; i++ {
		fmt.Print(u.underlineChar)
	}
	fmt.Println("")
	// ˄
}

// ˅

// ˄
