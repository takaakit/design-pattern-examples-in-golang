// ˅
package main

import "strconv"

// ˄

type FileSystemElement struct {
	// ˅

	// ˄

	name string

	size int

	// ˅

	// ˄
}

func NewFileSystemElement(name string, size int) *FileSystemElement {
	// ˅
	return &FileSystemElement{name, size}
	// ˄
}

func (self *FileSystemElement) ToString() string {
	// ˅
	return self.name + " (" + strconv.Itoa(self.size) + ")"
	// ˄
}

// ˅

// ˄
