// ˅
package main

import (
	"bytes"
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
	tablePage := &TablePage{}
	tablePage.Page = *NewPage(title, author)
	return tablePage
	// ˄
}

func (self *TablePage) ToHTML() string {
	// ˅
	var buffer bytes.Buffer
	buffer.WriteString("<html><head><title>" + self.title + "</title></head><body>\n")
	buffer.WriteString("<h1>" + self.title + "</h1>\n")
	buffer.WriteString("<table width=\"80%\" border=\"3\">\n")
	for i := 0; i < len(self.contents); i++ {
		buffer.WriteString("<tr>" + self.contents[i].ToHTML() + "</tr>\n")
	}
	buffer.WriteString("</table>\n")
	buffer.WriteString("<hr><address>" + self.author + "</address>\n")
	buffer.WriteString("</body></html>\n")
	return buffer.String()
	// ˄
}

// ˅

// ˄
