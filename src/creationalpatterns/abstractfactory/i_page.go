// ˅
package main

// ˄

type IPage interface {
	ToHTML() string

	Add(item Item)

	Output(context string)

	// ˅

	// ˄
}

// ˅

// ˄
