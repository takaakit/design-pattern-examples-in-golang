// ˅
package main

import (
	"math/rand"
)

// ˄

// When winning a game, show the same hand at the next time.
type StrategyA struct {
	// ˅

	// ˄

	won bool

	preHand *Hand

	// ˅

	// ˄
}

func NewStrategyA(randomSeed int64) *StrategyA {
	// ˅
	rand.Seed(randomSeed)

	strategyA := &StrategyA{}
	strategyA.won = false
	strategyA.preHand = getHand(rock)
	return strategyA
	// ˄
}

func (self *StrategyA) NextHand() *Hand {
	// ˅
	if self.won == false {
		self.preHand = getHand(rand.Intn(3))
	}
	return self.preHand
	// ˄
}

func (self *StrategyA) Learn(win bool) {
	// ˅
	self.won = win
	// ˄
}

// ˅

// ˄
