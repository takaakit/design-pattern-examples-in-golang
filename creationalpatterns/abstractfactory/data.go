// ˅
package main

// ˄

type Data struct {
	// ˅

	// ˄

	name string

	items []Item

	// ˅

	// ˄
}

func NewData(name string) *Data {
	// ˅
	data := &Data{}
	data.name = name
	return data
	// ˄
}

func (self *Data) Add(item Item) {
	// ˅
	self.items = append(self.items, item)
	// ˄
}

// ˅

// ˄
