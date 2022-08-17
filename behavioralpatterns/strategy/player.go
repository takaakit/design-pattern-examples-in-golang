// ˅
package strategy

import "strconv"

// ˄

type Player struct {
	// ˅

	// ˄

	name string

	winCount int

	lossCount int

	gameCount int

	strategy Strategy

	// ˅

	// ˄
}

func NewPlayer(name string, strategy Strategy) *Player {
	// ˅
	return &Player{name: name, winCount: 0, lossCount: 0, gameCount: 0, strategy: strategy}
	// ˄
}

// Show a hand signal from the strategy.
func (p *Player) ShowHandSignal() *HandSignal {
	// ˅
	return p.strategy.ShowHandSignal()
	// ˄
}

// Notify a game result.
func (p *Player) NotifyGameResult(result GameResultType, ownHand *HandSignal, opponentsHand *HandSignal) {
	// ˅
	switch result {
	case Win:
		p.winCount++
		p.gameCount++
	case Loss:
		p.lossCount++
		p.gameCount++
	case Draw:
		p.gameCount++
	}

	p.strategy.NotifyGameResult(result, ownHand, opponentsHand)
	// ˄
}

func (p *Player) String() string {
	// ˅
	return p.name + " [" + strconv.Itoa(p.gameCount) + " games, " + strconv.Itoa(p.winCount) + " won, " + strconv.Itoa(p.lossCount) + " lost, " + strconv.Itoa(p.gameCount-p.winCount-p.lossCount) + " drew]"
	// ˄
}

// ˅

// ˄
