// ˅
package main

// ˄

// Node in the syntax tree.
type INode interface {
	Parse(context *Context)

	ToString() string

	// ˅

	// ˄
}

// ˅

// ˄
