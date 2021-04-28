package main

// Factory to make ID cards.

func main() {
	factory := NewCreditCardFactory()
	var iFactory IFactory = factory

	jacksonCard := factory.Create(iFactory, "Jacson")
	jacksonCard.Use()

	sophiaCard := factory.Create(iFactory, "Sophia")
	sophiaCard.Use()

	oliviaCard := factory.Create(iFactory, "Olivia")
	oliviaCard.Use()
}
