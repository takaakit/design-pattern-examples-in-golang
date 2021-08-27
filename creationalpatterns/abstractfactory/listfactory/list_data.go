// ˅
package abstractfactory

import (
	"bytes"
	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/abstractfactory/factory"
)

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
	return &ListData{Data: *NewData(name)}
	// ˄
}

func (l *ListData) ToHTML() string {
	// ˅
	var buffer bytes.Buffer
	buffer.WriteString("<li>" + l.Name + "<ul>\n")
	for i := 0; i < len(l.Items); i++ {
		buffer.WriteString(l.Items[i].ToHTML())
	}
	buffer.WriteString("</ul></li>\n")
	return buffer.String()
	// ˄
}

// ˅

// ˄
