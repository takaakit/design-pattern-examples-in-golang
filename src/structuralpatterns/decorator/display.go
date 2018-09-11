// ˅
package main

import "fmt"

// ˄

type Display struct {
	// ˅

	// ˄

	// Column width
	columns int

	// Number of rows
	rows int

	// ˅

	// ˄
}

// Show all
func (self *Display) Show(iDisplay IDisplay) {
	// ˅
	for i := 0; i < iDisplay.GetRows(); i++ {
		fmt.Println(iDisplay.GetLineText(i))
	}
	// ˄
}

// ˅

// ˄
