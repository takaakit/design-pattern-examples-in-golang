// ˅
package observer

// ˄

// An abstract class that generates numbers.
type Subject struct {
	// ˅

	// ˄

	ISubject

	observers []Observer

	// ˅

	// ˄
}

func NewSubject() *Subject {
	// ˅
	return &Subject{observers: []Observer{}}
	// ˄
}

func (s *Subject) Attach(observer Observer) {
	// ˅
	s.observers = append(s.observers, observer)
	// ˄
}

func (s *Subject) Detach(observer Observer) {
	// ˅
	tmpObservers := []Observer{}
	for _, o := range s.observers {
		if o != observer {
			tmpObservers = append(tmpObservers, o)
		}
	}
	s.observers = tmpObservers
	// ˄
}

func (s *Subject) NotifyObservers() {
	// ˅
	for _, observer := range s.observers {
		observer.Update(s)
	}
	// ˄
}

// ˅

// ˄
