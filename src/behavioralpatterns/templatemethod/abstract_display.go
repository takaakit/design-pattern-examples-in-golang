// ˅
package main

// ˄

type AbstractDisplay struct {
	// ˅

	// ˄

	// ˅

	// ˄
}

func (self *AbstractDisplay) Output(iAbstractDisplay IAbstractDisplay) {
	// ˅
	iAbstractDisplay.Open()
	for i := 0; i < 5; i++ { // Repeat write 5 times
		iAbstractDisplay.Write()
	}
	iAbstractDisplay.Close()
	// ˄
}

// ˅

// ˄
