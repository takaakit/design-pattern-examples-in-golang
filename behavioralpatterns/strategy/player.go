// ˅
package main

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
	return &Player{name, 0, 0, 0, strategy}
	// ˄
}

// Calculate a hand from the strategy.
func (self *Player) NextHand() *Hand {
	// ˅
	return self.strategy.NextHand()
	// ˄
}

// Won a game.
func (self *Player) Won() {
	// ˅
	self.strategy.Learn(true)
	self.winCount++
	self.gameCount++
	// ˄
}

// Lost a game.
func (self *Player) Lost() {
	// ˅
	self.strategy.Learn(false)
	self.lossCount++
	self.gameCount++
	// ˄
}

// Drew a game.
func (self *Player) Drew() {
	// ˅
	self.gameCount++
	// ˄
}

func (self *Player) ToString() string {
	// ˅
	return self.name + " [" + strconv.Itoa(self.gameCount) + " games, " + strconv.Itoa(self.winCount) + " won, " + strconv.Itoa(self.lossCount) + " lost, " + strconv.Itoa(self.gameCount-self.winCount-self.lossCount) + " drew]"
	// ˄
}

// ˅

// ˄
