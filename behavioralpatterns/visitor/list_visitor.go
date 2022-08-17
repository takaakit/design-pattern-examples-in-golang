// ˅
package visitor

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
func (l *ListVisitor) VisitFile(file *File) {
	// ˅
	fmt.Println(l.currentDirectory + "/" + file.String())
	// ˄
}

// Visit a directory
func (l *ListVisitor) VisitDirectory(directory *Directory) {
	// ˅
	fmt.Println(l.currentDirectory + "/" + directory.String())
	var visitedDirectory = l.currentDirectory
	l.currentDirectory = l.currentDirectory + "/" + directory.name
	for _, iFileSystemElement := range directory.fileSystemElements {
		iFileSystemElement.Accept(l)
	}
	l.currentDirectory = visitedDirectory
	// ˄
}

// ˅

// ˄
