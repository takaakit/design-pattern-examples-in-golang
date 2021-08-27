// ˅
package bridge

// ˄

type MultiLineDisplay struct {
	// ˅

	// ˄

	Display

	// ˅

	// ˄
}

func NewMultiLineDisplay(impl DisplayImpl) *MultiLineDisplay {
	// ˅
	return &MultiLineDisplay{Display: *NewDisplay(impl)}
	// ˄
}

// Repeat display for the specified number of times
func (m *MultiLineDisplay) OutputMultiple(times int) {
	// ˅
	m.Open()
	for i := 0; i < times; i++ {
		m.Write()
	}
	m.Close()
	// ˄
}

// ˅

// ˄
