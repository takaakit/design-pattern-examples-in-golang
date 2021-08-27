// ˅
package interpreter

import "strconv"

// ˄

// A node corresponding to "repeat".
type Repeat struct {
	// ˅

	// ˄

	number int

	commandList Node

	// ˅

	// ˄
}

func NewRepeat() *Repeat {
	// ˅
	return &Repeat{number: 0}
	// ˄
}

func (r *Repeat) Parse(context *Context) {
	// ˅
	context.SlideToken("repeat")

	r.number = context.GetNumber()
	context.SlideToken(strconv.Itoa(r.number))

	aCommandList := NewCommandList()
	aCommandList.Parse(context)

	r.commandList = aCommandList // Hold the parsed command list
	// ˄
}

func (r *Repeat) String() string {
	// ˅
	return "repeat " + strconv.Itoa(r.number) + " " + r.commandList.String()
	// ˄
}

// ˅

// ˄
