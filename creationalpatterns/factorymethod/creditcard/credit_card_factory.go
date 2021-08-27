// ˅
package factorymethod

import (
	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/factorymethod/framework"
)

// ˄

type CreditCardFactory struct {
	// ˅

	// ˄

	Factory

	// ˅

	// ˄
}

func NewCreditCardFactory() *CreditCardFactory {
	// ˅
	return &CreditCardFactory{Factory: *NewFactory()}
	// ˄
}

func (c *CreditCardFactory) CreateProduct(owner string) Product {
	// ˅
	return NewCreditCard(owner)
	// ˄
}

// ˅

// ˄
