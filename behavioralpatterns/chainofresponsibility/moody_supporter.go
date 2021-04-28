// ˅
package main

// ˄

type MoodySupporter struct {
	// ˅

	// ˄

	Supporter

	// ˅

	// ˄
}

func NewMoodySupporter(name string) *MoodySupporter {
	// ˅
	return &MoodySupporter{Supporter{name: name}}
	// ˄
}

// Troubles with an odd ID are handled.
func (self *MoodySupporter) Handle(trouble *Trouble) bool {
	// ˅
	return trouble.id%2 == 1
	// ˄
}

// ˅

// ˄
