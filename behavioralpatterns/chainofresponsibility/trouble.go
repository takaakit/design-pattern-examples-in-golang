// ˅
package chainofresponsibility

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
	return &Trouble{id: id}
	// ˄
}

func (t *Trouble) String() string {
	// ˅
	return "[Trouble " + strconv.Itoa(t.id) + "]"
	// ˄
}

// ˅

// ˄
