// ˅
package main

// ˄

type Memento struct {
	// ˅

	// ˄

	// Money
	money int

	// Desserts
	desserts []string

	// ˅

	// ˄
}

func NewMemento(money int) *Memento {
	// ˅
	memento := &Memento{}
	memento.money = money
	return memento
	// ˄
}

// Add a dessert
func (self *Memento) addDessert(dessert string) {
	// ˅
	self.desserts = append(self.desserts, dessert)
	// ˄
}

// ˅

// ˄
