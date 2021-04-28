// ˅
package main

// ˄

type SpecialSupporter struct {
	// ˅

	// ˄

	Supporter

	targetId int

	// ˅

	// ˄
}

func NewSpecialSupporter(name string, targetId int) *SpecialSupporter {
	// ˅
	return &SpecialSupporter{Supporter{name: name}, targetId}
	// ˄
}

// Troubles with the specific ID are handled.
func (self *SpecialSupporter) Handle(trouble *Trouble) bool {
	// ˅
	return trouble.id == self.targetId
	// ˄
}

// ˅

// ˄
