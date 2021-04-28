// ˅
package main

// ˄

type Visitor interface {
	VisitFile(file *File)

	VisitDirectory(directory *Directory)

	// ˅

	// ˄
}

// ˅

// ˄
