// ˅
package abstractfactory

import (
	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/abstractfactory/factory"
)

// ˄

type TableLink struct {
	// ˅

	// ˄

	Link

	// ˅

	// ˄
}

func NewTableLink(name string, url string) *TableLink {
	// ˅
	return &TableLink{Link: *NewLink(name, url)}
	// ˄
}

func (t *TableLink) ToHTML() string {
	// ˅
	return "  <td><a href=\"" + t.Url + "\">" + t.Name + "</a></td>\n"
	// ˄
}

// ˅

// ˄
