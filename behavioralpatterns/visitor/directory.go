// ˅
package visitor

import "strconv"

// ˄

type Directory struct {
	// ˅

	// ˄

	name string

	// Collection of elements
	fileSystemElements []FileSystemElement

	// ˅

	// ˄
}

func NewDirectory(name string) *Directory {
	// ˅
	return &Directory{fileSystemElements: []FileSystemElement{}, name: name}
	// ˄
}

// Accept a visitor
func (d *Directory) Accept(visitor Visitor) {
	// ˅
	visitor.VisitDirectory(d)
	// ˄
}

func (d *Directory) GetName() string {
	// ˅
	return d.name
	// ˄
}

func (d *Directory) GetSize() int {
	// ˅
	var size int = 0
	for _, iFileSystemElement := range d.fileSystemElements {
		size += iFileSystemElement.GetSize()
	}
	return size
	// ˄
}

func (d *Directory) String() string {
	// ˅
	return d.GetName() + " (" + strconv.Itoa(int(d.GetSize())) + ")"
	// ˄
}

// Add an entry
func (d *Directory) Add(fileSystemElement FileSystemElement) {
	// ˅
	d.fileSystemElements = append(d.fileSystemElements, fileSystemElement)
	// ˄
}

// ˅

// ˄
