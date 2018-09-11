// ˅
package main

// ˄

type LimitedSupporter struct {
	// ˅

	// ˄

	Supporter

	limitId int

	// ˅

	// ˄
}

func NewLimitedSupporter(name string, limitId int) *LimitedSupporter {
	// ˅
	return &LimitedSupporter{Supporter{name: name}, limitId}
	// ˄
}

// Troubles with an ID below the limit are handled.
func (self *LimitedSupporter) Handle(trouble *Trouble) bool {
	// ˅
	return trouble.id <= self.limitId
	// ˄
}

// ˅

// ˄
