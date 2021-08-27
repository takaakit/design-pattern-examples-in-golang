// ˅
package abstractfactory

import (
	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/abstractfactory/factory"
)

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
	return &ListLink{Link: *NewLink(name, url)}
	// ˄
}

func (l *ListLink) ToHTML() string {
	// ˅
	return "  <li><a href=\"" + l.Url + "\">" + l.Name + "</a></li>\n"
	// ˄
}

// ˅

// ˄
