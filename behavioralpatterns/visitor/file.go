// ˅
package visitor

import "strconv"

// ˄

type File struct {
	// ˅

	// ˄

	name string

	size int32

	// ˅

	// ˄
}

func NewFile(name string, size int32) *File {
	// ˅
	return &File{name: name, size: size}
	// ˄
}

func (f *File) Accept(visitor Visitor) {
	// ˅
	visitor.VisitFile(f)
	// ˄
}

func (f *File) GetName() string {
	// ˅
	return f.name
	// ˄
}

func (f *File) GetSize() int32 {
	// ˅
	return f.size
	// ˄
}

func (f *File) String() string {
	// ˅
	return f.GetName() + " (" + strconv.Itoa(int(f.GetSize())) + ")"
	// ˄
}

// ˅

// ˄
