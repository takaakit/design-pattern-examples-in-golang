// ˅
package main

// ˄

type Command struct {
	// ˅

	// ˄

	node INode

	// ˅

	// ˄
}

func NewCommand() *Command {
	// ˅
	return &Command{}
	// ˄
}

func (self *Command) Parse(context *Context) {
	// ˅
	if context.GetToken() == "repeat" {
		self.node = NewRepeat()
	} else {
		self.node = NewAction(context.GetToken())
	}
	self.node.Parse(context)
	// ˄
}

func (self *Command) ToString() string {
	// ˅
	return self.node.ToString()
	// ˄
}

// ˅

// ˄
