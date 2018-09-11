// ˅
package main

// ˄

type ListLink struct {
	// ˅

	// ˄

	Link

	// ˅

	// ˄
}

func NewListLink(name string, url string) *ListLink {
	// ˅
	listLink := &ListLink{}
	listLink.Link = *NewLink(name, url)
	return listLink
	// ˄
}

func (self *ListLink) ToHTML() string {
	// ˅
	return "  <li><a href=\"" + self.url + "\">" + self.name + "</a></li>\n"
	// ˄
}

// ˅

// ˄
