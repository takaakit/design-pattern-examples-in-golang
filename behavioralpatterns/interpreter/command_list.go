// ˅
package interpreter

import (
	"fmt"
	"os"
)

// ˄

type CommandList struct {
	// ˅

	// ˄

	nodes []Node

	// ˅

	// ˄
}

func NewCommandList() *CommandList {
	// ˅
	return &CommandList{}
	// ˄
}

func (c *CommandList) Parse(context *Context) {
	// ˅
	for {
		if context.GetToken() == "" {
			fmt.Println("Missing 'end'")
			os.Exit(1)
		} else if context.GetToken() == "end" {
			context.SlideToken("end")
			break
		} else {
			aNode := NewCommand()
			aNode.Parse(context)

			c.nodes = append(c.nodes, aNode) // Hold the parsed node
		}
	}
	// ˄
}

func (c *CommandList) String() string {
	// ˅
	str := "["
	for i, n := range c.nodes {
		str += n.String()
		if i < len(c.nodes)-1 {
			str += ", "
		}
	}
	str += "]"
	return str
	// ˄
}

// ˅

// ˄
