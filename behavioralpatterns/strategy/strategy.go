// ˅
package strategy

// ˄

type Strategy interface {
	ShowHandSignal() *HandSignal

	NotifyGameResult(result GameResultType, ownHand *HandSignal, opponentsHand *HandSignal)

	// ˅

	// ˄
}

// ˅

// ˄
