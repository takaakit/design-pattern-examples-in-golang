// ˅
package bridge

// ˄

type Display struct {
	// ˅

	// ˄

	impl DisplayImpl

	// ˅

	// ˄
}

func NewDisplay(impl DisplayImpl) *Display {
	// ˅
	return &Display{impl: impl}
	// ˄
}

func (d *Display) Output() {
	// ˅
	d.Open()
	d.Write()
	d.Close()
	// ˄
}

func (d *Display) Open() {
	// ˅
	d.impl.ImplOpen()
	// ˄
}

func (d *Display) Write() {
	// ˅
	d.impl.ImplWrite()
	// ˄
}

func (d *Display) Close() {
	// ˅
	d.impl.ImplClose()
	// ˄
}

// ˅

// ˄
