// ˅
package bridge

import (
	"fmt"
	"strconv"
)

// ˄

type Directory struct {
	// ˅

	// ˄

	name string

	elements []FileSystemElement

	// ˅

	// ˄
}

func NewDirectory(name string) *Directory {
	// ˅
	return &Directory{name: name, elements: []FileSystemElement{}}
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
	for i := 0; i < len(d.elements); i++ {
		size += d.elements[i].GetSize()
	}
	return size
	// ˄
}

// Print this element with the "upperPath".
func (d *Directory) Print(upperPath string) {
	// ˅
	fmt.Println(upperPath + "/" + d.String())
	for i := 0; i < len(d.elements); i++ {
		d.elements[i].Print(upperPath + "/" + d.name)
	}
	// ˄
}

func (d *Directory) String() string {
	// ˅
	return d.GetName() + " (" + strconv.Itoa(d.GetSize()) + ")"
	// ˄
}

// Add an element
func (d *Directory) Add(element FileSystemElement) {
	// ˅
	d.elements = append(d.elements, element)
	// ˄
}

// ˅

// ˄
