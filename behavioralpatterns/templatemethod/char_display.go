// ˅
package main

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
	charDisplay := &CharDisplay{}
	charDisplay.chText = chText
	return charDisplay
	// ˄
}

func (self *CharDisplay) Open() {
	// ˅
	fmt.Print("<<") // Display "<<" in the start characters.
	// ˄
}

func (self *CharDisplay) Write() {
	// ˅
	fmt.Print(self.chText) // Display the character.
	// ˄
}

func (self *CharDisplay) Close() {
	// ˅
	fmt.Println(">>") // Display ">>" in the end characters.
	// ˄
}

// ˅

// ˄
