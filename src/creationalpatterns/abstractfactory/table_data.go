// ˅
package main

import (
	"bytes"
	"strconv"
)

// ˄

type TableData struct {
	// ˅

	// ˄

	Data

	// ˅

	// ˄
}

func NewTableData(name string) *TableData {
	// ˅
	tableData := &TableData{}
	tableData.Data = *NewData(name)
	return tableData
	// ˄
}

func (self *TableData) ToHTML() string {
	// ˅
	var buffer bytes.Buffer
	buffer.WriteString("<td><table width=\"100%\" border=\"2\">\n")
	buffer.WriteString("<tr><td bgcolor=\"#00CC00\" align=\"center\" colspan=\"" + strconv.Itoa(len(self.items)) + "\"><b>" + self.name + "</b></td></tr>\n")
	buffer.WriteString("<tr>\n")
	for i := 0; i < len(self.items); i++ {
		buffer.WriteString(self.items[i].ToHTML())
	}
	buffer.WriteString("</tr>\n")
	buffer.WriteString("</table></td>\n")
	return buffer.String()
	// ˄
}

// ˅

// ˄
