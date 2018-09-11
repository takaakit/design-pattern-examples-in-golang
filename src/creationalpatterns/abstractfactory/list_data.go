// ˅
package main

import "bytes"

// ˄

type ListData struct {
	// ˅

	// ˄

	Data

	// ˅

	// ˄
}

func NewListData(name string) *ListData {
	// ˅
	listData := &ListData{}
	listData.Data = *NewData(name)
	return listData
	// ˄
}

func (self *ListData) ToHTML() string {
	// ˅
	var buffer bytes.Buffer
	buffer.WriteString("<li>" + self.name + "<ul>\n")
	for i := 0; i < len(self.items); i++ {
		buffer.WriteString(self.items[i].ToHTML())
	}
	buffer.WriteString("</ul></li>\n")
	return buffer.String()
	// ˄
}

// ˅

// ˄
