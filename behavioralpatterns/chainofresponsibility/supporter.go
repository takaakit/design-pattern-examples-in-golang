// ˅
package chainofresponsibility

import (
	"fmt"
)

// ˄

type Supporter struct {
	// ˅

	// ˄

	ISupporter

	// Supporter name
	name string

	// Next supporter
	Next ISupporter

	// ˅

	// ˄
}

func NewSupporter(name string) *Supporter {
	// ˅
	return &Supporter{name: name, Next: nil}
	// ˄
}

// Trouble support
// Troubles are sent around.
// 
// Client-Specified Self pattern.
func (s *Supporter) Support(iSupporter ISupporter, trouble *Trouble) {
	// ˅
	if iSupporter.CanHandle(trouble) {
		s.supported(trouble)
	} else if s.Next != nil {
		s.Next.Support(s.Next, trouble)
	} else {
		s.unsupported(trouble)
	}
	// ˄
}

func (s *Supporter) String() string {
	// ˅
	return "[" + s.name + "]"
	// ˄
}

// Trouble was supported.
func (s *Supporter) supported(trouble *Trouble) {
	// ˅
	fmt.Println(trouble.String() + " was handled by " + s.String() + ".")
	// ˄
}

// Trouble was unsupported.
func (s *Supporter) unsupported(trouble *Trouble) {
	// ˅
	fmt.Println(trouble.String() + " was not handled.")
	// ˄
}

// ˅

// ˄
