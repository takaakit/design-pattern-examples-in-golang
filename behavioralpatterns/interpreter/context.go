// ˅
package interpreter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ˄

// Analyze the syntax
type Context struct {
	// ˅

	// ˄

	tokens []string

	currentIndex int

	// ˅

	// ˄
}

func NewContext(text string) *Context {
	// ˅
	return &Context{strings.Fields(text), 0}
	// ˄
}

func (c *Context) NextToken() string {
	// ˅
	c.currentIndex++
	if c.currentIndex < len(c.tokens) {
		return c.tokens[c.currentIndex]
	} else {
		return ""
	}
	// ˄
}

func (c *Context) GetToken() string {
	// ˅
	return c.tokens[c.currentIndex]
	// ˄
}

func (c *Context) SlideToken(token string) {
	// ˅
	if token != c.GetToken() {
		fmt.Println("WARNING: " + token + " is expected but " + c.GetToken() + " was found.")
		os.Exit(1)
	}
	c.NextToken()
	// ˄
}

func (c *Context) GetNumber() int {
	// ˅
	i, err := strconv.Atoi(c.GetToken())
	if err != nil {
		fmt.Println("WARNING: " + c.GetToken())
		os.Exit(1)
	}
	return i
	// ˄
}

// ˅

// ˄
