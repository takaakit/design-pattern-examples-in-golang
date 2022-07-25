// ˅
package main

// ˄

// Adapt the MessageDisplay interface to the Print interface.
type PrintMessageDisplay struct {
	// ˅

	// ˄

	MessageDisplay

	// ˅

	// ˄
}

func NewPrintMessageDisplay(message string) *PrintMessageDisplay {
	// ˅
	return &PrintMessageDisplay{MessageDisplay: *NewMessageDisplay(message)}
	// ˄
}

func (p *PrintMessageDisplay) PrintWeak() {
	// ˅
	p.DisplayWithHyphens()
	// ˄
}

func (p *PrintMessageDisplay) PrintStrong() {
	// ˅
	p.DisplayWithBrackets()
	// ˄
}

// ˅

// ˄
