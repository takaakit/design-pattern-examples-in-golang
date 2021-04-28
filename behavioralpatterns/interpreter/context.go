// ˅
package main

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

	nodes []string

	currentIndex int

	// ˅

	// ˄
}

func NewContext(text string) *Context {
	// ˅
	return &Context{strings.Fields(text), 0}
	// ˄
}

func (self *Context) NextToken() string {
	// ˅
	self.currentIndex++
	if self.currentIndex < len(self.nodes) {
		return self.nodes[self.currentIndex]
	} else {
		return ""
	}
	// ˄
}

func (self *Context) GetToken() string {
	// ˅
	return self.nodes[self.currentIndex]
	// ˄
}

func (self *Context) SlideToken(token string) {
	// ˅
	if token != self.GetToken() {
		fmt.Println("WARNING: " + token + " is expected but " + self.GetToken() + " was found.")
		os.Exit(1)
	}
	self.NextToken()
	// ˄
}

func (self *Context) GetNumber() int {
	// ˅
	i, err := strconv.Atoi(self.GetToken())
	if err != nil {
		fmt.Println("WARNING: " + self.GetToken())
		os.Exit(1)
	}
	return i
	// ˄
}

// ˅

// ˄
