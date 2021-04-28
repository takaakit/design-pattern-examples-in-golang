// ˅
package main

import "fmt"

// ˄

type MessageDisplay struct {
	// ˅

	// ˄

	message string

	// ˅

	// ˄
}

func NewMessageDisplay(message string) *MessageDisplay {
	// ˅
	return &MessageDisplay{message}
	// ˄
}

func (self *MessageDisplay) DisplayWithHyphens() {
	// ˅
	fmt.Println("-- " + self.message + " --")
	// ˄
}

func (self *MessageDisplay) DisplayWithBrackets() {
	// ˅
	fmt.Println("[[ " + self.message + " ]]")
	// ˄
}

// ˅

// ˄
