// ˅
package interpreter

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

func NewAction() *Action {
	// ˅
	return &Action{name: ""}
	// ˄
}

func (a *Action) Parse(context *Context) {
	// ˅
	currentToken := context.GetToken()
	if currentToken != "forward" && currentToken != "right" && currentToken != "left" {
		fmt.Println(currentToken + " is unknown")
		os.Exit(1)
	}

	a.name = currentToken // Hold the current token as this action name

	context.SlideToken(currentToken)
	// ˄
}

func (a *Action) String() string {
	// ˅
	return a.name
	// ˄
}

// ˅

// ˄
