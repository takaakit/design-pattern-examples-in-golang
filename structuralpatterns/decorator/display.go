// ˅
package decorator

// ˄

type Display interface {
	// Column width
	GetColumns() int

	// Number of rows
	GetRows() int

	GetLineText(row int) string

	// Show all
	Show()

	// ˅

	// ˄
}

// ˅

// ˄
