// ˅
package mediator

// ˄

type Colleague struct {
	// ˅

	// ˄

	IColleague

	mediator Mediator

	// ˅

	// ˄
}

func NewColleague() *Colleague {
	// ˅
	return &Colleague{}
	// ˄
}

func (c *Colleague) SetMediator(mediator Mediator) {
	// ˅
	c.mediator = mediator
	// ˄
}

// ˅

// ˄
