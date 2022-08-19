// ˅
package mediator

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
	return &ColleagueTextField{Colleague: *NewColleague(), lineEdit: lineEdit}
	// ˄
}

// Set enable/disable from the Mediator
func (c *ColleagueTextField) SetActivation(isEnable bool) {
	// ˅
	_ = c.lineEdit.SetReadOnly(!isEnable)
	// ˄
}

func (c *ColleagueTextField) OnTextChanged() {
	// ˅
	c.mediator.ColleagueChanged()
	// ˄
}

func (c *ColleagueTextField) IsEmpty() bool {
	// ˅
	return c.lineEdit.Text() == ""
	// ˄
}

// ˅

// ˄
