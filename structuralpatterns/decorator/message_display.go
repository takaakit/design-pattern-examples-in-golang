// ˅
package decorator

import "fmt"

// ˄

type MessageDisplay struct {
	// ˅

	// ˄

	// Message to be displayed
	message string

	// ˅

	// ˄
}

func NewMessageDisplay(message string) *MessageDisplay {
	// ˅
	return &MessageDisplay{message: message}
	// ˄
}

func (m *MessageDisplay) GetLineText(row int) string {
	// ˅
	if row == 0 {
		return m.message
	} else {
		return ""
	}
	// ˄
}

func (m *MessageDisplay) GetColumns() int {
	// ˅
	return len(m.message)
	// ˄
}

func (m *MessageDisplay) GetRows() int {
	// ˅
	return 1
	// ˄
}

// Show all
func (m *MessageDisplay) Show() {
	// ˅
	for i := 0; i < m.GetRows(); i++ {
		fmt.Println(m.GetLineText(i))
	}
	// ˄
}

// ˅

// ˄
