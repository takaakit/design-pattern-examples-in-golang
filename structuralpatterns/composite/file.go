// ˅
package bridge

import (
	"fmt"
	"strconv"
)

// ˄

type File struct {
	// ˅

	// ˄

	name string

	size int

	// ˅

	// ˄
}

func NewFile(name string, size int) *File {
	// ˅
	return &File{name: name, size: size}
	// ˄
}

func (f *File) GetName() string {
	// ˅
	return f.name
	// ˄
}

func (f *File) GetSize() int {
	// ˅
	return f.size
	// ˄
}

// Print this element with the "upperPath".
func (f *File) Print(upperPath string) {
	// ˅
	fmt.Println(upperPath + "/" + f.String())
	// ˄
}

func (f *File) String() string {
	// ˅
	return f.GetName() + " (" + strconv.Itoa(f.GetSize()) + ")"
	// ˄
}

// ˅

// ˄
