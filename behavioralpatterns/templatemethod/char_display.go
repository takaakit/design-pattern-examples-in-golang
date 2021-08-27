// ˅
package templatemethod

import (
	"fmt"
)

// ˄

type CharDisplay struct {
	// ˅

	// ˄

	AbstractDisplay

	chText string

	// ˅

	// ˄
}

func NewCharDisplay(chText string) *CharDisplay {
	// ˅
	return &CharDisplay{AbstractDisplay: AbstractDisplay{}, chText: chText}
	// ˄
}

func (c *CharDisplay) Open() {
	// ˅
	fmt.Print("<<") // Display "<<" in the start characters.
	// ˄
}

func (c *CharDisplay) Write() {
	// ˅
	fmt.Print(c.chText) // Display the character.
	// ˄
}

func (c *CharDisplay) Close() {
	// ˅
	fmt.Println(">>") // Display ">>" in the end characters.
	// ˄
}

// ˅

// ˄
