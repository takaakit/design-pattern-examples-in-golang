package main

// Display only one line or display the specified number of lines.

func main() {
	d1 := NewDisplay(NewTextDisplayImpl("Japan"))
	d2 := NewMultiLineDisplay(NewTextDisplayImpl("The United States of America"))
	d3 := NewMultiLineDisplay(NewTextDisplayImpl("Brazil"))
	d1.Output()
	d2.Output()
	d3.Output()
	d3.OutputMultiple(3)
}
