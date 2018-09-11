// ˅
package main

// ˄

type CreditCardFactory struct {
	// ˅

	// ˄

	Factory

	cardOwners []string

	// ˅

	// ˄
}

func NewCreditCardFactory() *CreditCardFactory {
	// ˅
	creditCardFactory := &CreditCardFactory{}
	creditCardFactory.Factory = *NewFactory()
	return creditCardFactory
	// ˄
}

func (self *CreditCardFactory) CreateProduct(owner string) Product {
	// ˅
	return NewCreditCard(owner)
	// ˄
}

func (self *CreditCardFactory) RegisterProduct(product Product) {
	// ˅
	self.cardOwners = append(self.cardOwners, product.(*CreditCard).owner)
	// ˄
}

func (self *CreditCardFactory) GetCardOwner() []string {
	// ˅
	return self.cardOwners
	// ˄
}

// ˅

// ˄
