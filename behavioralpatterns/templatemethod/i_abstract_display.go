// ˅
package templatemethod

// ˄

type IAbstractDisplay interface {
	Open()

	Write()

	Close()

	Output(iAbstractDisplay IAbstractDisplay)

	// ˅

	// ˄
}

// ˅

// ˄
