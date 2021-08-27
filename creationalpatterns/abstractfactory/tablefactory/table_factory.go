// ˅
package abstractfactory

import (
	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/abstractfactory/factory"
)

// ˄

type TableFactory struct {
	// ˅

	// ˄

	// ˅

	// ˄
}

func NewTableFactory() *TableFactory {
	// ˅
	return &TableFactory{}
	// ˄
}

func (t *TableFactory) CreatePage(title string, author string) IPage {
	// ˅
	return NewTablePage(title, author)
	// ˄
}

func (t *TableFactory) CreateLink(name string, url string) ILink {
	// ˅
	return NewTableLink(name, url)
	// ˄
}

func (t *TableFactory) CreateData(name string) IData {
	// ˅
	return NewTableData(name)
	// ˄
}

// ˅

// ˄
