// ˅
package factory

// ˄

type IPage interface {
	ToHTML() string

	Add(item Item)

	// Note: This is the client-specified self pattern.
	Output(iPage IPage)

	// ˅

	// ˄
}

// ˅

// ˄
