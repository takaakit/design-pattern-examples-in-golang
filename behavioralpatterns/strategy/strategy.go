// ˅
package main

// ˄

type Strategy interface {
	NextHand() *Hand

	Learn(win bool)

	// ˅

	// ˄
}

// ˅

// ˄
