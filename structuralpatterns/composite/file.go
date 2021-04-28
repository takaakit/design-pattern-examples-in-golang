// ˅
package main

import "fmt"

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

// Print this element with the "upperPath".
func (self *File) Print(upperPath string) {
	// ˅
	fmt.Println(upperPath + "/" + self.ToString())
	// ˄
}

// ˅

// ˄
