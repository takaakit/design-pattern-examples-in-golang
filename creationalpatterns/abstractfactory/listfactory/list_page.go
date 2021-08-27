// ˅
package abstractfactory

import (
	"bytes"
	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/abstractfactory/factory"
)

// ˄

type ListPage struct {
	// ˅

	// ˄

	Page

	// ˅

	// ˄
}

func NewListPage(title string, author string) *ListPage {
	// ˅
	return &ListPage{Page: *NewPage(title, author)}
	// ˄
}

func (l *ListPage) ToHTML() string {
	// ˅
	var buffer bytes.Buffer
	buffer.WriteString("<html><head><title>" + l.Title + "</title></head>\n")
	buffer.WriteString("<body><h1>" + l.Title + "</h1>\n")
	buffer.WriteString("<ul>\n")
	for i := 0; i < len(l.Contents); i++ {
		buffer.WriteString(l.Contents[i].ToHTML())
	}
	buffer.WriteString("</ul>\n")
	buffer.WriteString("<hr><address>" + l.Author + "</address>")
	buffer.WriteString("</body></html>\n")
	return buffer.String()
	// ˄
}

// ˅

// ˄
