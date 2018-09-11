// ˅
package main

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
	colleagueRadioButton := &ColleagueRadioButton{}
	colleagueRadioButton.Colleague = *NewColleague()
	colleagueRadioButton.radioButton = radioButton
	return colleagueRadioButton
	// ˄
}

// Set enable/disable from the Mediator
func (self *ColleagueRadioButton) SetActivation(isEnable bool) {
	// ˅
	self.radioButton.SetEnabled(isEnable)
	// ˄
}

func (self *ColleagueRadioButton) OnClicked() {
	// ˅
	self.mediator.ColleagueChanged()
	// ˄
}

func (self *ColleagueRadioButton) IsSelected() bool {
	// ˅
	return self.radioButton.Checked()
	// ˄
}

// ˅

// ˄
