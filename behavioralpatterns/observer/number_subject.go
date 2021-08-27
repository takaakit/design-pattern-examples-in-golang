// ˅
package observer

// ˄

// Generate a random number.
type NumberSubject struct {
	// ˅

	// ˄

	Subject

	value int

	// ˅

	// ˄
}

func NewNumberSubject() *NumberSubject {
	// ˅
	return &NumberSubject{Subject: *NewSubject(), value: 0}
	// ˄
}

func (n *NumberSubject) SetValue(value int) {
	// ˅
	n.value = value
	n.NotifyObservers()
	// ˄
}

func (n *NumberSubject) GetValue() int {
	// ˅
	return n.value
	// ˄
}

// ˅

// ˄
