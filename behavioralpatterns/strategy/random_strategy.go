// ˅
package strategy

import (
	"math/rand"
	"time"
)

// ˄

// Random Strategy: showing a random hand signal.
type RandomStrategy struct {
	// ˅

	// ˄

	// ˅

	// ˄
}

func NewRandomStrategy() *RandomStrategy {
	// ˅
	rand.Seed(time.Now().UnixNano())

	return &RandomStrategy{}
	// ˄
}

func (r *RandomStrategy) ShowHandSignal() *HandSignal {
	// ˅
	switch rand.Intn(3) {
	case 0:
		return getHand(Rock)
	case 1:
		return getHand(Scissors)
	default:
		return getHand(Paper)
	}
	// ˄
}

func (r *RandomStrategy) NotifyGameResult(result GameResultType, ownHand *HandSignal, opponentsHand *HandSignal) {
	// ˅
	// Do nothing
	// ˄
}

// ˅

// ˄
