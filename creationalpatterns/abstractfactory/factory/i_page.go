// ˅
package factory

// ˄

type IPage interface {
	ToHTML() string

	Add(item Item)

	// Client-Specified Self pattern
	Output(iPage IPage)

	// ˅

	// ˄
}

// ˅

// ˄
