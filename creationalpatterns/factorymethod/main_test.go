package factorymethod

import (
	"testing"

	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/factorymethod/creditcard"
)

/*
The subject is a factory to make credit cards. The Factory defines how to create a credit card,
but the actual credit card is created by the CreditCardFactory.
The "createProduct()" is called a Factory Method, and it is responsible for manufacturing an object.
*/

func TestMain(m *testing.M) {
	factory := NewCreditCardFactory()

	jacksonCard := factory.Create(factory, "Jackson")
	jacksonCard.Use()

	sophiaCard := factory.Create(factory, "Sophia")
	sophiaCard.Use()

	oliviaCard := factory.Create(factory, "Olivia")
	oliviaCard.Use()
}
