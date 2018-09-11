// ˅
package main

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
	tableLink := &TableLink{}
	tableLink.Link = *NewLink(name, url)
	return tableLink
	// ˄
}

func (self *TableLink) ToHTML() string {
	// ˅
	return "  <td><a href=\"" + self.url + "\">" + self.name + "</a></td>\n"
	// ˄
}

// ˅

// ˄
