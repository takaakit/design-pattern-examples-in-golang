// ˅
package decorator

// ˄

type Frame struct {
	// ˅

	// ˄

	display Display

	// ˅

	// ˄
}

func NewFrame(display Display) *Frame {
	// ˅
	return &Frame{display: display}
	// ˄
}

// ˅

// ˄
