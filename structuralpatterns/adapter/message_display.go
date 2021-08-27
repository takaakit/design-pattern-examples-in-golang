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
	return &MessageDisplay{message: message}
	// ˄
}

func (m *MessageDisplay) DisplayWithHyphens() {
	// ˅
	fmt.Println("-- " + m.message + " --")
	// ˄
}

func (m *MessageDisplay) DisplayWithBrackets() {
	// ˅
	fmt.Println("[[ " + m.message + " ]]")
	// ˄
}

// ˅

// ˄
