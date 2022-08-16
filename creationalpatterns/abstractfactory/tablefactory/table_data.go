// ˅
package abstractfactory

import (
	"bytes"
	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/abstractfactory/factory"
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
	return &TableData{Data: *NewData(name)}
	// ˄
}

func (t *TableData) ToHTML() string {
	// ˅
	var buffer bytes.Buffer
	buffer.WriteString("<td><table width=\"100%\" border=\"2\">\n")
	buffer.WriteString("<tr><td bgcolor=\"#00CC00\" align=\"center\" colspan=\"" + strconv.Itoa(len(t.Items)) + "\"><b>" + t.Name + "</b></td></tr>\n")
	buffer.WriteString("<tr>\n")
	for _, item := range t.Items {
		buffer.WriteString(item.ToHTML())
	}
	buffer.WriteString("</tr>\n")
	buffer.WriteString("</table></td>\n")
	return buffer.String()
	// ˄
}

// ˅

// ˄
