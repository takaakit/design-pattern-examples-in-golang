// ˅
package main

import "strconv"

// ˄

// A node corresponding to "repeat".
type Repeat struct {
	// ˅

	// ˄

	number int

	commandList INode

	// ˅

	// ˄
}

func NewRepeat() *Repeat {
	// ˅
	return &Repeat{number: 0}
	// ˄
}

func (self *Repeat) Parse(context *Context) {
	// ˅
	context.SlideToken("repeat")
	self.number = context.GetNumber()
	context.NextToken()
	self.commandList = NewCommandList()
	self.commandList.Parse(context)
	// ˄
}

func (self *Repeat) ToString() string {
	// ˅
	return "[repeat " + strconv.Itoa(self.number) + " " + self.commandList.ToString() + "]"
	// ˄
}

// ˅

// ˄
