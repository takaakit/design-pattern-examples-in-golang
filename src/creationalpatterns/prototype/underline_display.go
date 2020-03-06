// ˅
package main

import "fmt"

// ˄

type UnderlineDisplay struct {
	// ˅

	// ˄

	underlineChar string

	// ˅

	// ˄
}

func NewUnderlineDisplay(underlineChar string) *UnderlineDisplay {
	// ˅
	return &UnderlineDisplay{underlineChar}
	// ˄
}

func (self *UnderlineDisplay) CreateClone() Display {
	// ˅
	return NewUnderlineDisplay(self.underlineChar)
	// ˄
}

func (self *UnderlineDisplay) Show(message string) {
	// ˅
	var length = len(message)
	fmt.Println("\"" + message + "\"")
	fmt.Print(" ")
	for i := 0; i < length; i++ {
		fmt.Print(self.underlineChar)
	}
	fmt.Println("")
	// ˄
}

// ˅

// ˄
