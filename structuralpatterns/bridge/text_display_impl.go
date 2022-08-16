// ˅
package bridge

import (
	"fmt"
	"strings"
)

// ˄

type TextDisplayImpl struct {
	// ˅

	// ˄

	// A string to display
	text string

	// A number of characters in bytes
	width int

	// ˅

	// ˄
}

func NewTextDisplayImpl(text string) *TextDisplayImpl {
	// ˅
	return &TextDisplayImpl{
		text:  text,
		width: len(text), // Set the number of characters in bytes.
	}
	// ˄
}

func (t *TextDisplayImpl) ImplOpen() {
	// ˅
	t.printLine()
	// ˄
}

func (t *TextDisplayImpl) ImplWrite() {
	// ˅
	fmt.Println(":" + t.text + ":") // Enclose a text with ":" and display it.
	// ˄
}

func (t *TextDisplayImpl) ImplClose() {
	// ˅
	t.printLine()
	// ˄
}

func (t *TextDisplayImpl) printLine() {
	// ˅
	fmt.Print("*")                          // Display "*" mark at the beginning of a frame.
	fmt.Print(strings.Repeat(".", t.width)) // Display "." for the number of "width".
	fmt.Println("*") // Display "*" mark at the end of a frame.
	// ˄
}

// ˅

// ˄
