// ˅
package main

import (
	"bytes"
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
	listPage := &ListPage{}
	listPage.Page = *NewPage(title, author)
	return listPage
	// ˄
}

func (self *ListPage) ToHTML() string {
	// ˅
	var buffer bytes.Buffer
	buffer.WriteString("<html><head><title>" + self.title + "</title></head>\n")
	buffer.WriteString("<body><h1>" + self.title + "</h1>\n")
	buffer.WriteString("<ul>\n")
	for i := 0; i < len(self.contents); i++ {
		buffer.WriteString(self.contents[i].ToHTML())
	}
	buffer.WriteString("</ul>\n")
	buffer.WriteString("<hr><address>" + self.author + "</address>")
	buffer.WriteString("</body></html>\n")
	return buffer.String()
	// ˄
}

// ˅

// ˄
