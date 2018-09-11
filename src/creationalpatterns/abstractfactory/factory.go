// ˅
package main

// ˄

type Factory interface {
	CreatePage(title string, author string) IPage

	CreateLink(name string, url string) ILink

	CreateData(name string) IData

	// ˅

	// ˄
}

// ˅

// ˄
