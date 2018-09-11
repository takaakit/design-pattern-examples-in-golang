// ˅
package main

import "fmt"

// ˄

type Directory struct {
	// ˅

	// ˄

	FileSystemElement

	elements []IFileSystemElement

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

// Print this element with the "upperPath".
func (self *Directory) Print(upperPath string) {
	// ˅
	fmt.Println(upperPath + "/" + self.ToString())
	for i := 0; i < len(self.elements); i++ {
		self.elements[i].Print(upperPath + "/" + self.name)
	}
	// ˄
}

// Add a element
func (self *Directory) Add(element IFileSystemElement) IFileSystemElement {
	// ˅
	self.elements = append(self.elements, element)
	return element
	// ˄
}

// ˅

// ˄
