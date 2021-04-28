// ˅
package main

import (
	"fmt"
)

// ˄

type ListVisitor struct {
	// ˅

	// ˄

	// Currently visited directory
	currentDirectory string

	// ˅

	// ˄
}

func NewListVisitor() *ListVisitor {
	// ˅
	return &ListVisitor{""}
	// ˄
}

// Visit a file
func (self *ListVisitor) VisitFile(file *File) {
	// ˅
	fmt.Println(self.currentDirectory + "/" + file.ToString())
	// ˄
}

// Visit a directory
func (self *ListVisitor) VisitDirectory(directory *Directory) {
	// ˅
	fmt.Println(self.currentDirectory + "/" + directory.ToString())
	var visitedDirectory = self.currentDirectory
	self.currentDirectory = self.currentDirectory + "/" + directory.name
	for _, fileSystemElement := range directory.elements {
		var element Element = fileSystemElement
		element.Accept(self)
	}
	self.currentDirectory = visitedDirectory
	// ˄
}

// ˅

// ˄
