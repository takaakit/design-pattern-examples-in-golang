// ˅
package observer

// ˄

// Provides an interface for attaching and detaching Observer objects.
type ISubject interface {
	Attach(observer Observer)

	Detach(observer Observer)

	NotifyObservers()

	// ˅

	// ˄
}

// ˅

// ˄
