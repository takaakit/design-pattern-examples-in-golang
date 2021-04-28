// ˅
package main

import (
	"fmt"
	"os"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// ˄

type AppLogin struct {
	// ˅

	// ˄

	radioLogin *ColleagueRadioButton

	radioGuest *ColleagueRadioButton

	textUsername *ColleagueTextField

	textPassword *ColleagueTextField

	buttonOk *ColleagueButton

	buttonCancel *ColleagueButton

	// ˅

	// ˄
}

func NewAppLogin() *AppLogin {
	// ˅
	appLogin := &AppLogin{}
	appLogin.CreateColleagues()

	return appLogin
	// ˄
}

func (self *AppLogin) CreateColleagues() {
	// ˅
	// Create LineEdit, PushButton and RadioButton

	var rd1 *walk.RadioButton
	var rd2 *walk.RadioButton
	var edit1 *walk.LineEdit
	var edit2 *walk.LineEdit
	var pb1 *walk.PushButton
	var pb2 *walk.PushButton

	self.radioGuest = NewColleagueRadioButton(rd1)
	self.radioLogin = NewColleagueRadioButton(rd2)
	self.textUsername = NewColleagueTextField(edit1)
	self.textPassword = NewColleagueTextField(edit2)
	self.buttonOk = NewColleagueButton(pb1)
	self.buttonCancel = NewColleagueButton(pb2)

	// Set mediators
	self.radioLogin.mediator = self
	self.radioGuest.mediator = self
	self.textUsername.mediator = self
	self.textPassword.mediator = self
	self.buttonOk.mediator = self
	self.buttonCancel.mediator = self

	// Create main window
	mainWindow := MainWindow{
		Title:  "Mediator Example",
		Size:   Size{250, 200},
		Layout: VBox{},
		Children: []Widget{
			RadioButtonGroupBox{
				Layout: HBox{},
				Buttons: []RadioButton{
					RadioButton{
						AssignTo:  &self.radioGuest.radioButton,
						Text:      "Guest",
						OnClicked: self.radioGuest.OnClicked,
					},
					RadioButton{
						AssignTo:  &self.radioLogin.radioButton,
						Text:      "Login",
						OnClicked: self.radioLogin.OnClicked,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Text: "Username:",
					},
					LineEdit{
						AssignTo:      &self.textUsername.lineEdit,
						MinSize:       Size{110, 0},
						OnTextChanged: self.textUsername.OnTextChanged,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Text: "Password:",
					},
					LineEdit{
						AssignTo:      &self.textPassword.lineEdit,
						MinSize:       Size{110, 0},
						PasswordMode:  true,
						OnTextChanged: self.textPassword.OnTextChanged,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						AssignTo:  &self.buttonOk.pushButton,
						Text:      "OK",
						OnClicked: self.buttonOk.OnClicked,
					},
					PushButton{
						AssignTo:  &self.buttonCancel.pushButton,
						Text:      "Cancel",
						OnClicked: self.buttonCancel.OnClicked,
					},
				},
			},
		},
	}

	if _, err := mainWindow.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// ˄
}

// Change enable/disable of the Colleagues when notified from the Mediators.
func (self *AppLogin) ColleagueChanged() {
	// ˅
	if self.buttonOk.IsSelected() == true || self.buttonCancel.IsSelected() == true {
		os.Exit(0)
	} else {
		if self.radioGuest.IsSelected() == true { // Guest mode
			self.textUsername.SetActivation(false)
			self.textPassword.SetActivation(false)
			self.buttonOk.SetActivation(true)
		} else { // Login mode
			self.textUsername.SetActivation(true)
			self.textPassword.SetActivation(true)

			// Judge whether the changed Colleage is enabled or disabled
			if self.textUsername.IsEmpty() == false && self.textPassword.IsEmpty() == false {
				self.buttonOk.SetActivation(true)
			} else {
				self.buttonOk.SetActivation(false)
			}
		}
	}
	// ˄
}

// ˅

// ˄
