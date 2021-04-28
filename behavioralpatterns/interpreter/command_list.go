// ˅
package main

import (
	"fmt"
	"os"
)

// ˄

type CommandList struct {
	// ˅

	// ˄

	nodes []INode

	// ˅

	// ˄
}

func NewCommandList() *CommandList {
	// ˅
	return &CommandList{}
	// ˄
}

func (self *CommandList) Parse(context *Context) {
	// ˅
	for {
		if context.GetToken() == "" {
			fmt.Println("Missing 'end'")
			os.Exit(1)
		} else if context.GetToken() == "end" {
			context.SlideToken("end")
			break
		} else {
			commandNode := NewCommand()
			commandNode.Parse(context)
			self.nodes = append(self.nodes, commandNode)
		}
	}
	// ˄
}

func (self *CommandList) ToString() string {
	// ˅
	var str string
	for i := 0; i < len(self.nodes); i++ {
		str += (self.nodes[i]).ToString()
		if i < len(self.nodes)-1 {
			str += " "
		}
	}
	return str
	// ˄
}

// ˅

// ˄
