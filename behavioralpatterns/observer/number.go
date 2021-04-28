// ˅
package main

// ˄

// An abstract class that generates numbers.
type Number struct {
	// ˅

	// ˄

	value int

	observers []Observer

	// ˅

	// ˄
}

func NewNumber(value int) *Number {
	// ˅
	number := &Number{}
	number.value = value
	return number
	// ˄
}

func (self *Number) AddObserver(observer Observer) {
	// ˅
	self.observers = append(self.observers, observer)
	// ˄
}

func (self *Number) DeleteObserver(observer Observer) {
	// ˅
	retObservers := []Observer{}
	for _, o := range self.observers {
		if o != observer {
			retObservers = append(retObservers, o)
		}
	}
	self.observers = retObservers
	// ˄
}

func (self *Number) NotifyObservers() {
	// ˅
	for _, observer := range self.observers {
		observer.Update(self)
	}
	// ˄
}

// ˅

// ˄
