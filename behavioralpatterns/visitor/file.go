// ˅
package main

// ˄

type File struct {
	// ˅

	// ˄

	FileSystemElement

	// ˅

	// ˄
}

func NewFile(name string, size int) *File {
	// ˅
	file := &File{}
	file.FileSystemElement = *NewFileSystemElement(name, size)
	return file
	// ˄
}

func (self *File) Accept(visitor Visitor) {
	// ˅
	visitor.VisitFile(self)
	// ˄
}

// ˅

// ˄
