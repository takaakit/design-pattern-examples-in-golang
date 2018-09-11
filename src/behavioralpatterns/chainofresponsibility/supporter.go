// ˅
package main

import (
	"fmt"
)

// ˄

type Supporter struct {
	// ˅

	// ˄

	// Supporter name
	name string

	// Next supporter
	next ISupporter

	// ˅

	// ˄
}

func NewSupporter(name string) *Supporter {
	// ˅
	supporter := new(Supporter)
	supporter.name = name
	supporter.next = nil
	return supporter
	// ˄
}

// Trouble support
// Troubles are sent around.
func (self *Supporter) Support(iSupporter ISupporter, trouble *Trouble) {
	// ˅
	if iSupporter.Handle(trouble) {
		self.supported(trouble)
	} else if self.next != nil {
		self.next.Support(self.next, trouble)
	} else {
		self.unsupported(trouble)
	}
	// ˄
}

// Set a next supporter.
func (self *Supporter) SetNext(next ISupporter) ISupporter {
	// ˅
	self.next = next
	return next
	// ˄
}

func (self *Supporter) ToString() string {
	// ˅
	return "[" + self.name + "]"
	// ˄
}

// Trouble was supported.
func (self *Supporter) supported(trouble *Trouble) {
	// ˅
	fmt.Println(trouble.ToString() + " was handled by " + self.ToString() + ".")
	// ˄
}

// Trouble was unsupported.
func (self *Supporter) unsupported(trouble *Trouble) {
	// ˅
	fmt.Println(trouble.ToString() + " was not handled.")
	// ˄
}

// ˅

// ˄
