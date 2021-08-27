// ˅
package strategy

// Hands of rock-scissors-paper
var handSignals = map[HandSignalType]*HandSignal{
	Rock:NewHand(Rock),
	Scissors:NewHand(Scissors),
	Paper:NewHand(Paper),
}

// Characters of the hands
var name = map[HandSignalType]string{
	Rock:"Rock",
	Scissors:"Scissors",
	Paper:"Paper",
}

// Get an instance of the hand
func getHand(handValue HandSignalType) *HandSignal {
	return handSignals[handValue]
}

// ˄

type HandSignal struct {
	// ˅

	// ˄

	// Values of rock, scissors and paper.
	value HandSignalType

	// ˅

	// ˄
}

func NewHand(value HandSignalType) *HandSignal {
	// ˅
	return &HandSignal{value: value}
	// ˄
}

// Return true if "this" is stronger than "hand".
func (h *HandSignal) IsStrongerThan(hand *HandSignal) bool {
	// ˅
	return h.judgeGame(hand) == 1
	// ˄
}

// Return false if "this" is weaker than "hand".
func (h *HandSignal) IsWeakerThan(hand *HandSignal) bool {
	// ˅
	return h.judgeGame(hand) == -1
	// ˄
}

func (h *HandSignal) ToString() string {
	// ˅
	return name[h.value]
	// ˄
}

// The draw is 0. "this" win is 1. "hand" win is -1.
func (h *HandSignal) judgeGame(hand *HandSignal) int {
	// ˅
	if h.value == hand.value {
		return 0
	} else if (h.value+1)%3 == hand.value {
		return 1
	} else {
		return -1
	}
	// ˄
}

// ˅

// ˄
