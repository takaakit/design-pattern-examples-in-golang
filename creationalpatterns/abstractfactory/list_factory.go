// ˅
package main

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

func (self *ListFactory) CreatePage(title string, author string) IPage {
	// ˅
	return NewListPage(title, author)
	// ˄
}

func (self *ListFactory) CreateLink(name string, url string) ILink {
	// ˅
	return NewListLink(name, url)
	// ˄
}

func (self *ListFactory) CreateData(name string) IData {
	// ˅
	return NewListData(name)
	// ˄
}

// ˅

// ˄
