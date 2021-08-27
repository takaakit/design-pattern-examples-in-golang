// ˅
package factorymethod

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

	return &CreditCard{owner: owner}
	// ˄
}

func (c *CreditCard) Use() {
	// ˅
	fmt.Println("Use " + c.owner + "'s card.")
	// ˄
}

// ˅

// ˄
