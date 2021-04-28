// ˅
package main

// Hands of rock-scissors-paper
const (
	rock     = 0 // Rock
	scissors = 1 // Scissors
	paper    = 2 // Paper
)

// Characters of the hands
var hands = []*Hand{NewHand(rock), NewHand(scissors), NewHand(paper)}

// Characters of the hands
var name = []string{"Rock", "Scissors", "Paper"}

// Get an instance of the hand
func getHand(handValue int) *Hand {
	return hands[handValue]
}

// ˄

type Hand struct {
	// ˅

	// ˄

	// Values of rock, scissors and paper.
	value int

	// ˅

	// ˄
}

func NewHand(value int) *Hand {
	// ˅
	return &Hand{value}
	// ˄
}

// Return true if "this" is stronger than "hand".
func (self *Hand) IsStrongerThan(hand *Hand) bool {
	// ˅
	return self.judgeGame(hand) == 1
	// ˄
}

// Return false if "this" is weaker than "hand".
func (self *Hand) IsWeakerThan(hand *Hand) bool {
	// ˅
	return self.judgeGame(hand) == -1
	// ˄
}

func (self *Hand) ToString() string {
	// ˅
	return name[self.value]
	// ˄
}

// The draw is 0. "this" win is 1. "hand" win is -1.
func (self *Hand) judgeGame(hand *Hand) int {
	// ˅
	if self.value == hand.value {
		return 0
	} else if (self.value+1)%3 == hand.value {
		return 1
	} else {
		return -1
	}
	// ˄
}

// ˅

// ˄
