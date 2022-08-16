// ˅
package templatemethod

import (
	"fmt"
	"strings"
)

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
	return &StringDisplay{AbstractDisplay: AbstractDisplay{}, strText: strText, width: len(strText)}
	// ˄
}

func (s *StringDisplay) Open() {
	// ˅
	s.writeLine() // Write a line
	// ˄
}

func (s *StringDisplay) Write() {
	// ˅
	fmt.Println("|" + s.strText + "|") // Display the character with "|"
	// ˄
}

func (s *StringDisplay) Close() {
	// ˅
	s.writeLine() // Write a line
	// ˄
}

func (s *StringDisplay) writeLine() {
	// ˅
	fmt.Print("+")                          // Display an end mark "+"
	fmt.Print(strings.Repeat("-", s.width)) // Display a line "-"
	fmt.Println("+")                        //  Display an end mark "+"
	// ˄
}

// ˅

// ˄
