// ˅
package interpreter

// ˄

// Node in the syntax tree.
type Node interface {
	Parse(context *Context)

	String() string

	// ˅

	// ˄
}

// ˅

// ˄
