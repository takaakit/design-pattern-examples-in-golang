// ˅
package mediator

import (
	"github.com/lxn/walk"
)

// ˄

type ColleagueButton struct {
	// ˅

	// ˄

	Colleague

	pushButton *walk.PushButton

	// ˅

	// ˄
}

func NewColleagueButton(pushButton *walk.PushButton) *ColleagueButton {
	// ˅
	return &ColleagueButton{Colleague: *NewColleague(), pushButton: pushButton}
	// ˄
}

// Set enable/disable from the Mediator
func (c *ColleagueButton) SetActivation(isEnable bool) {
	// ˅
	c.pushButton.SetEnabled(isEnable)
	// ˄
}

func (c *ColleagueButton) OnClicked() {
	// ˅
	c.mediator.ColleagueChanged()
	// ˄
}

func (c *ColleagueButton) IsPressed() bool {
	// ˅
	return c.pushButton.Focused()
	// ˄
}

// ˅

// ˄
