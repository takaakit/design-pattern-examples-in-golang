// ˅
package main

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
	return &LazySupporter{Supporter{name: name}}
	// ˄
}

// No troubles are handled.
func (self *LazySupporter) Handle(trouble *Trouble) bool {
	// ˅
	return false
	// ˄
}

// ˅

// ˄
