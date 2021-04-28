// ˅
package main

// ˄

type Singleton struct {
	// ˅

	// ˄

	// ˅

	// ˄
}

func NewSingleton() *Singleton {
	// ˅
	if instance != nil {
		instance = &Singleton{}
	}
	return instance
	// ˄
}

// ˅
var instance *Singleton = nil

// ˄
