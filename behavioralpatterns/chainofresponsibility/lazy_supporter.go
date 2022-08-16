// ˅
package chainofresponsibility

// ˄

type LazySupporter struct {
	// ˅

	// ˄

	Supporter

	// ˅

	// ˄
}

func NewLazySupporter(name string) *LazySupporter {
	// ˅
	return &LazySupporter{Supporter: *NewSupporter(name)}
	// ˄
}

// No troubles are handled.
func (l *LazySupporter) CanHandle(trouble *Trouble) bool {
	// ˅
	return false
	// ˄
}

// ˅

// ˄
