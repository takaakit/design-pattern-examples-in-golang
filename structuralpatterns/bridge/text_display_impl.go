// ˅
package main

import "fmt"

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
	textDisplayImpl := &TextDisplayImpl{}
	textDisplayImpl.text = text
	textDisplayImpl.width = len(text)
	return textDisplayImpl
	// ˄
}

func (self *TextDisplayImpl) ImplOpen() {
	// ˅
	self.printLine()
	// ˄
}

func (self *TextDisplayImpl) ImplWrite() {
	// ˅
	fmt.Println(":" + self.text + ":") // Enclose a text with ":" and display it.
	// ˄
}

func (self *TextDisplayImpl) ImplClose() {
	// ˅
	self.printLine()
	// ˄
}

func (self *TextDisplayImpl) printLine() {
	// ˅
	fmt.Print("*")                    // Display "*" mark at the beginning of a frame.
	for i := 0; i < self.width; i++ { // Display "." for the number of "width".
		fmt.Print(".")
	}
	fmt.Println("*") // Display "*" mark at the end of a frame.
	// ˄
}

// ˅

// ˄
