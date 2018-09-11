// ˅
package main

import (
	"math/rand"
	"time"
)

// ˄

// Generate a random number.
type RandomNumber struct {
	// ˅

	// ˄

	Number

	// ˅

	// ˄
}

func NewRandomNumber() *RandomNumber {
	// ˅
	rand.Seed(time.Now().UnixNano())
	return &RandomNumber{}
	// ˄
}

func (self *RandomNumber) Generate() {
	// ˅
	for i := 0; i < 20; i++ {
		self.value = rand.Intn(50)
		self.NotifyObservers()
	}
	// ˄
}

// ˅

// ˄
