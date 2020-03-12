// ˅
package main

import "fmt"

// ˄

type CreditCard struct {
	// ˅

	// ˄

	owner string

	// ˅

	// ˄
}

func NewCreditCard(owner string) *CreditCard {
	// ˅
	fmt.Println("Make " + owner + "'s card.")

	creditCard := &CreditCard{}
	creditCard.owner = owner
	return creditCard
	// ˄
}

func (self *CreditCard) Use() {
	// ˅
	fmt.Println("Use " + self.owner + "'s card.")
	// ˄
}

// ˅

// ˄
