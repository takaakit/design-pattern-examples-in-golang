// ˅
package chainofresponsibility

// ˄

type SpecialSupporter struct {
	// ˅

	// ˄

	Supporter

	targetID int

	// ˅

	// ˄
}

func NewSpecialSupporter(name string, targetID int) *SpecialSupporter {
	// ˅
	return &SpecialSupporter{Supporter: Supporter{name: name}, targetID: targetID}
	// ˄
}

// Troubles with the specific ID are handled.
func (s *SpecialSupporter) CanHandle(trouble *Trouble) bool {
	// ˅
	return trouble.id == s.targetID
	// ˄
}

// ˅

// ˄
