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
	for _, item := range l.Items {
		buffer.WriteString(item.ToHTML())
	}
	buffer.WriteString("</ul></li>\n")
	return buffer.String()
	// ˄
}

// ˅

// ˄
