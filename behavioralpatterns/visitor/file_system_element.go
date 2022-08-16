// ˅
package visitor

// ˄

type FileSystemElement interface {
	Element

	GetName() string

	GetSize() int

	String() string

	// ˅

	// ˄
}

// ˅

// ˄
