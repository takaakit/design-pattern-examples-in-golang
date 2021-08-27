// ˅
package visitor

// ˄

type Visitor interface {
	VisitFile(file *File)

	VisitDirectory(directory *Directory)

	// ˅

	// ˄
}

// ˅

// ˄
