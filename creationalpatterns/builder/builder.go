// ˅
package main

// ˄

type Builder interface {
	CreateTitle(title string)

	CreateSection(section string)

	CreateItems(items []string)

	Close()

	// ˅

	// ˄
}

// ˅

// ˄
