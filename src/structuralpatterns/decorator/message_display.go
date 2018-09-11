// ˅
package main

// ˄

type MessageDisplay struct {
	// ˅

	// ˄

	Display

	// Message to be displayed
	message string

	// ˅

	// ˄
}

func NewMessageDisplay(message string) *MessageDisplay {
	// ˅
	messageDisplay := &MessageDisplay{}
	messageDisplay.message = message
	return messageDisplay
	// ˄
}

func (self *MessageDisplay) GetLineText(row int) string {
	// ˅
	if row == 0 {
		return self.message
	} else {
		return ""
	}
	// ˄
}

func (self *MessageDisplay) GetColumns() int {
	// ˅
	return len(self.message)
	// ˄
}

func (self *MessageDisplay) GetRows() int {
	// ˅
	return 1
	// ˄
}

// ˅

// ˄
