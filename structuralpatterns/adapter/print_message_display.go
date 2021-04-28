// ˅
package main

// ˄

type PrintMessageDisplay struct {
	// ˅

	// ˄

	MessageDisplay

	// ˅

	// ˄
}

func NewPrintMessageDisplay(message string) *PrintMessageDisplay {
	// ˅
	printMessageDisplay := &PrintMessageDisplay{}
	printMessageDisplay.MessageDisplay = *NewMessageDisplay(message)
	return printMessageDisplay
	// ˄
}

func (self *PrintMessageDisplay) PrintWeak() {
	// ˅
	self.DisplayWithHyphens()
	// ˄
}

func (self *PrintMessageDisplay) PrintStrong() {
	// ˅
	self.DisplayWithBrackets()
	// ˄
}

// ˅

// ˄
