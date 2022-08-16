// ˅
package chainofresponsibility

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
	return &LimitedSupporter{Supporter: *NewSupporter(name), limitId: limitId}
	// ˄
}

// Troubles with an ID below the limit are handled.
func (l *LimitedSupporter) CanHandle(trouble *Trouble) bool {
	// ˅
	return trouble.id <= l.limitId
	// ˄
}

// ˅

// ˄
