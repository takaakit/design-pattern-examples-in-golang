// ˅
package main

// ˄

type Directory struct {
	// ˅

	// ˄

	FileSystemElement

	// Collection of elements
	elements []Element

	// ˅

	// ˄
}

func NewDirectory(name string) *Directory {
	// ˅
	directory := &Directory{}
	directory.FileSystemElement = *NewFileSystemElement(name, 0)
	return directory
	// ˄
}

// Accept a visitor
func (self *Directory) Accept(visitor Visitor) {
	// ˅
	visitor.VisitDirectory(self)
	// ˄
}

// Add an entry
func (self *Directory) Add(element Element) *Directory {
	// ˅
	self.elements = append(self.elements, element)
	return self
	// ˄
}

// ˅

// ˄
