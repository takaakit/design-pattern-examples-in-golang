// ˅
package chainofresponsibility

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
	return &MoodySupporter{Supporter: *NewSupporter(name)}
	// ˄
}

// Troubles with an odd ID are handled.
func (m *MoodySupporter) CanHandle(trouble *Trouble) bool {
	// ˅
	return trouble.id%2 == 1
	// ˄
}

// ˅

// ˄
