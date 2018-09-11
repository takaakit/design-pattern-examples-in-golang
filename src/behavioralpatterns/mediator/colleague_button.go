// ˅
package main

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
	colleagueButton := &ColleagueButton{}
	colleagueButton.Colleague = *NewColleague()
	colleagueButton.pushButton = pushButton
	return colleagueButton
	// ˄
}

// Set enable/disable from the Mediator
func (self *ColleagueButton) SetActivation(isEnable bool) {
	// ˅
	self.pushButton.SetEnabled(isEnable)
	// ˄
}

func (self *ColleagueButton) OnClicked() {
	// ˅
	self.mediator.ColleagueChanged()
	// ˄
}

func (self *ColleagueButton) IsSelected() bool {
	// ˅
	return self.pushButton.Focused()
	// ˄
}

// ˅

// ˄
