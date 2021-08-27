// ˅
package interpreter

// ˄

type Command struct {
	// ˅

	// ˄

	node Node

	// ˅

	// ˄
}

func NewCommand() *Command {
	// ˅
	return &Command{}
	// ˄
}

func (c *Command) Parse(context *Context) {
	// ˅
	var aNode Node
	if context.GetToken() == "repeat" {
		aNode = NewRepeat()
		aNode.Parse(context)
	} else {
		aNode = NewAction()
		aNode.Parse(context)
	}

	c.node = aNode // Hold the parsed node
	// ˄
}

func (c *Command) String() string {
	// ˅
	return c.node.String()
	// ˄
}

// ˅

// ˄
