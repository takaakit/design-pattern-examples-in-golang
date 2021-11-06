// ˅
package chainofresponsibility

// ˄

type ISupporter interface {
	CanHandle(trouble *Trouble) bool

	// Trouble support
	// Troubles are sent around.
	// Note: This is the client-specified self pattern.
	Support(iSupporter ISupporter, trouble *Trouble)

	// ˅

	// ˄
}

// ˅

// ˄
