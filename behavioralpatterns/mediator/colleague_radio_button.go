// ˅
package mediator

import (
	"github.com/lxn/walk"
)

// ˄

type ColleagueRadioButton struct {
	// ˅

	// ˄

	Colleague

	radioButton *walk.RadioButton

	// ˅

	// ˄
}

func NewColleagueRadioButton(radioButton *walk.RadioButton) *ColleagueRadioButton {
	// ˅
	return &ColleagueRadioButton{Colleague: *NewColleague(), radioButton: radioButton}
	// ˄
}

// Set enable/disable from the Mediator
func (c *ColleagueRadioButton) SetActivation(isEnable bool) {
	// ˅
	c.radioButton.SetEnabled(isEnable)
	// ˄
}

func (c *ColleagueRadioButton) OnClicked() {
	// ˅
	c.mediator.ColleagueChanged()
	// ˄
}

func (c *ColleagueRadioButton) IsSelected() bool {
	// ˅
	return c.radioButton.Checked()
	// ˄
}

// ˅

// ˄
