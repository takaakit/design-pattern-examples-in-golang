// ˅
package main

import (
	"fmt"
	"os"
)

// ˄

// A node corresponding to "forward", "left", and "right".
type Action struct {
	// ˅

	// ˄

	name string

	// ˅

	// ˄
}

func NewAction(name string) *Action {
	// ˅
	return &Action{name: name}
	// ˄
}

func (self *Action) Parse(context *Context) {
	// ˅
	self.name = context.GetToken()
	context.SlideToken(self.name)
	if self.name != "forward" && self.name != "right" && self.name != "left" {
		fmt.Println(self.name + " is unknown")
		os.Exit(1)
	}
	// ˄
}

func (self *Action) ToString() string {
	// ˅
	return self.name
	// ˄
}

// ˅

// ˄
