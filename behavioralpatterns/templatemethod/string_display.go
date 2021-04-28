// ˅
package main

import "fmt"

// ˄

type StringDisplay struct {
	// ˅

	// ˄

	AbstractDisplay

	strText string

	// String width
	width int

	// ˅

	// ˄
}

func NewStringDisplay(strText string) *StringDisplay {
	// ˅
	stringDisplay := &StringDisplay{}
	stringDisplay.strText = strText
	stringDisplay.width = len(strText)
	return stringDisplay
	// ˄
}

func (self *StringDisplay) Open() {
	// ˅
	self.writeLine() // Write a line
	// ˄
}

func (self *StringDisplay) Write() {
	// ˅
	fmt.Println("|" + self.strText + "|") // Display the character with "|"
	// ˄
}

func (self *StringDisplay) Close() {
	// ˅
	self.writeLine() // Write a line
	// ˄
}

func (self *StringDisplay) writeLine() {
	// ˅
	fmt.Print("+") // Display an end mark "+"
	for i := 0; i < self.width; i++ {
		fmt.Print("-") // Display a line "-"
	}
	fmt.Println("+") //  Display an end mark "+"
	// ˄
}

// ˅

// ˄
