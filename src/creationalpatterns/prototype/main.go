package main

// Display a character string enclosed with a frame line, or drawn with an underline.

func main() {
	// Create a manager
	manager := NewManager()
	emphasisUnderline := NewUnderlineDisplay("~")
	highlightFrame := NewFrameDisplay("+")
	warningFrame := NewFrameDisplay("#")
	manager.RegisterDisplay("emphasis", emphasisUnderline)
	manager.RegisterDisplay("highlight", highlightFrame)
	manager.RegisterDisplay("warning", warningFrame)

	// Create displays
	display1 := manager.GetDisplay("emphasis")
	display1.Show("Nice to meet you.")
	display2 := manager.GetDisplay("highlight")
	display2.Show("Nice to meet you.")
	display1 = manager.GetDisplay("warning")
	display1.Show("Nice to meet you.")
}
