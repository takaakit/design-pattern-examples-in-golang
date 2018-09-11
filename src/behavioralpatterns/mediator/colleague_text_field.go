// ˅
package main

import (
	"github.com/lxn/walk"
)

// ˄

type ColleagueTextField struct {
	// ˅

	// ˄

	Colleague

	lineEdit *walk.LineEdit

	// ˅

	// ˄
}

func NewColleagueTextField(lineEdit *walk.LineEdit) *ColleagueTextField {
	// ˅
	colleagueTextField := &ColleagueTextField{}
	colleagueTextField.Colleague = *NewColleague()
	colleagueTextField.lineEdit = lineEdit
	return colleagueTextField
	// ˄
}

// Set enable/disable from the Mediator
func (self *ColleagueTextField) SetActivation(isEnable bool) {
	// ˅
	self.lineEdit.SetReadOnly(!isEnable)
	// ˄
}

func (self *ColleagueTextField) OnTextChanged() {
	// ˅
	self.mediator.ColleagueChanged()
	// ˄
}

func (self *ColleagueTextField) IsEmpty() bool {
	// ˅
	return self.lineEdit.Text() == ""
	// ˄
}

// ˅

// ˄
