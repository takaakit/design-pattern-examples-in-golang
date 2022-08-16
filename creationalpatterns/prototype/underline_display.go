// ˅
package prototype

import (
	"fmt"
	"strings"

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
	fmt.Println("\"" + message + "\"")
	fmt.Println(" " + strings.Repeat(u.underlineChar, len(message)))
	// ˄
}

// ˅

// ˄
