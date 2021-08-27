// ˅
package strategy

// ˄

// Mirror Strategy: showing a hand signal from the previous opponent's hand signal.
type MirrorStrategy struct {
	// ˅

	// ˄

	preOpponentsHand *HandSignal

	// ˅

	// ˄
}

func NewMirrorStrategy() *MirrorStrategy {
	// ˅
	return &MirrorStrategy{preOpponentsHand: getHand(Rock)}
	// ˄
}

func (m *MirrorStrategy) ShowHandSignal() *HandSignal {
	// ˅
	return m.preOpponentsHand
	// ˄
}

func (m *MirrorStrategy) NotifyGameResult(result GameResultType, ownHand *HandSignal, opponentsHand *HandSignal) {
	// ˅
	m.preOpponentsHand = opponentsHand
	// ˄
}

// ˅

// ˄
