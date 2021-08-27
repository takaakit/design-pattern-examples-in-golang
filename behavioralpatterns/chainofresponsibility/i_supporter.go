// ˅
package chainofresponsibility

// ˄

type ISupporter interface {
	CanHandle(trouble *Trouble) bool

	// Trouble support
	// Troubles are sent around.
	// 
	// Client-Specified Self pattern.
	Support(iSupporter ISupporter, trouble *Trouble)

	// ˅

	// ˄
}

// ˅

// ˄
