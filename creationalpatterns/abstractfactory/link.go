// ˅
package main

// ˄

type Link struct {
	// ˅

	// ˄

	name string

	url string

	// ˅

	// ˄
}

func NewLink(name string, url string) *Link {
	// ˅
	link := &Link{}
	link.name = name
	link.url = url
	return link
	// ˄
}

// ˅

// ˄
