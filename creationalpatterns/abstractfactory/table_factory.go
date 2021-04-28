// ˅
package main

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

func (self *TableFactory) CreatePage(title string, author string) IPage {
	// ˅
	return NewTablePage(title, author)
	// ˄
}

func (self *TableFactory) CreateLink(name string, url string) ILink {
	// ˅
	return NewTableLink(name, url)
	// ˄
}

func (self *TableFactory) CreateData(name string) IData {
	// ˅
	return NewTableData(name)
	// ˄
}

// ˅

// ˄
