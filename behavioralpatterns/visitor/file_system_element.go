// ˅
package visitor

// ˄

type FileSystemElement interface {
	Element

	GetName() string

	GetSize() int32

	String() string

	// ˅

	// ˄
}

// ˅

// ˄
