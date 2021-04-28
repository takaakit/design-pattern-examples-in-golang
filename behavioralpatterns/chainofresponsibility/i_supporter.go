// ˅
package main

// ˄

type ISupporter interface {
	SetNext(next ISupporter) ISupporter

	Handle(trouble *Trouble) bool

	Support(iSupporter ISupporter, trouble *Trouble)

	// ˅

	// ˄
}

// ˅

// ˄
