// ˅
package interpreter

// ˄

// A node corresponding to "program".
type Program struct {
	// ˅

	// ˄

	commandList Node

	// ˅

	// ˄
}

func NewProgram() *Program {
	// ˅
	return &Program{}
	// ˄
}

func (p *Program) Parse(context *Context) {
	// ˅
	context.SlideToken("program")

	aCommandList := NewCommandList()
	aCommandList.Parse(context)

	p.commandList = aCommandList // Hold the parsed command list
	// ˄
}

func (p *Program) String() string {
	// ˅
	return "[program " + p.commandList.String() + "]"
	// ˄
}

// ˅

// ˄
