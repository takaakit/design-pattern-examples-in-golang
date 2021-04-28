// ˅
package main

import (
	"math/rand"
)

// ˄

// Calculate a hand from the previous hand stochastically.
type StrategyB struct {
	// ˅

	// ˄

	history [][]int

	preHand *Hand

	curHand *Hand

	// ˅

	// ˄
}

func NewStrategyB(randomSeed int64) *StrategyB {
	// ˅
	rand.Seed(randomSeed)

	strategyB := &StrategyB{}
	strategyB.history = [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	strategyB.preHand = getHand(rock)
	strategyB.curHand = getHand(rock)
	return strategyB
	// ˄
}

func (self *StrategyB) NextHand() *Hand {
	// ˅
	var randomHandValue = rand.Intn(self.getHistorySum(self.curHand.value))
	var handValue = 0
	if randomHandValue < self.history[self.curHand.value][0] {
		handValue = rock
	} else if randomHandValue < self.history[self.curHand.value][0]+self.history[self.curHand.value][1] {
		handValue = scissors
	} else {
		handValue = paper
	}
	self.preHand = self.curHand
	self.curHand = getHand(handValue)
	return self.curHand
	// ˄
}

func (self *StrategyB) Learn(win bool) {
	// ˅
	if win {
		self.history[self.preHand.value][self.curHand.value]++
	} else {
		self.history[self.preHand.value][(self.curHand.value+1)%3]++
		self.history[self.preHand.value][(self.curHand.value+2)%3]++
	}
	// ˄
}

func (self *StrategyB) getHistorySum(handValue int) int {
	// ˅
	var sum = 0
	for i := 0; i < 3; i++ {
		sum += self.history[handValue][i]
	}
	return sum
	// ˄
}

// ˅

// ˄
