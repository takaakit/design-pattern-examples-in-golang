// ˅
package main

// ˄

// A node corresponding to "program".
type Program struct {
	// ˅

	// ˄

	commandList INode

	// ˅

	// ˄
}

func NewProgram() *Program {
	// ˅
	return &Program{}
	// ˄
}

func (self *Program) Parse(context *Context) {
	// ˅
	context.SlideToken("program")
	self.commandList = NewCommandList()
	self.commandList.Parse(context)
	// ˄
}

func (self *Program) ToString() string {
	// ˅
	if self.commandList != nil {
		return "[program " + self.commandList.ToString() + "]"
	} else {
		return ""
	}
	// ˄
}

// ˅

// ˄
