// ˅
package main

import "strconv"

// ˄

type Trouble struct {
	// ˅

	// ˄

	// Trouble number
	id int

	// ˅

	// ˄
}

func NewTrouble(id int) *Trouble {
	// ˅
	trouble := &Trouble{id: id}
	return trouble
	// ˄
}

func (self *Trouble) ToString() string {
	// ˅
	return "[Trouble " + strconv.Itoa(self.id) + "]"
	// ˄
}

// ˅

// ˄
