// ˅
package abstractfactory

import (
	"bytes"
	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/abstractfactory/factory"
)

// ˄

type TablePage struct {
	// ˅

	// ˄

	Page

	// ˅

	// ˄
}

func NewTablePage(title string, author string) *TablePage {
	// ˅
	return &TablePage{Page: *NewPage(title, author)}
	// ˄
}

func (t *TablePage) ToHTML() string {
	// ˅
	var buffer bytes.Buffer
	buffer.WriteString("<html><head><title>" + t.Title + "</title></head><body>\n")
	buffer.WriteString("<h1>" + t.Title + "</h1>\n")
	buffer.WriteString("<table width=\"80%\" border=\"3\">\n")
	for _, content := range t.Contents {
		buffer.WriteString("<tr>" + content.ToHTML() + "</tr>\n")
	}
	buffer.WriteString("</table>\n")
	buffer.WriteString("<hr><address>" + t.Author + "</address>\n")
	buffer.WriteString("</body></html>\n")
	return buffer.String()
	// ˄
}

// ˅

// ˄
