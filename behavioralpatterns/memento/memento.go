// ˅
package memento

// ˄

type Memento struct {
	// ˅

	// ˄

	// Money
	money int

	// ˅

	// ˄
}

func NewMemento(money int) *Memento {
	// ˅
	return &Memento{money: money}
	// ˄
}

func (m *Memento) Money() int {
	// ˅
	return m.money
	// ˄
}

// ˅

// ˄
