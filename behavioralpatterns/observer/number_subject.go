// ˅
package observer

// ˄

// Holds a value and notifies observers when the value is set.
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
	// Notify observers when the value is set.
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
