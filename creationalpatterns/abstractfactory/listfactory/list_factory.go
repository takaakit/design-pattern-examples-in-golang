// ˅
package abstractfactory

import (
	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/abstractfactory/factory"
)

// ˄

type ListFactory struct {
	// ˅

	// ˄

	// ˅

	// ˄
}

func NewListFactory() *ListFactory {
	// ˅
	return &ListFactory{}
	// ˄
}

func (l *ListFactory) CreatePage(title string, author string) IPage {
	// ˅
	return NewListPage(title, author)
	// ˄
}

func (l *ListFactory) CreateLink(name string, url string) ILink {
	// ˅
	return NewListLink(name, url)
	// ˄
}

func (l *ListFactory) CreateData(name string) IData {
	// ˅
	return NewListData(name)
	// ˄
}

// ˅

// ˄
