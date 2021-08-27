// ˅
package bridge

// ˄

type FileSystemElement interface {
	GetName() string

	GetSize() int

	// Print this element with the "upperPath".
	Print(upperPath string)

	String() string

	// ˅

	// ˄
}

// ˅

// ˄
