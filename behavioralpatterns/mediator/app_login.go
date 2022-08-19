// ˅
package mediator

import (
	"fmt"
	"os"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
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

func (a *AppLogin) CreateColleagues() {
	// ˅
	// Create LineEdit, PushButton and RadioButton
	var rd1 *walk.RadioButton
	var rd2 *walk.RadioButton
	var edit1 *walk.LineEdit
	var edit2 *walk.LineEdit
	var pb1 *walk.PushButton
	var pb2 *walk.PushButton

	a.radioGuest = NewColleagueRadioButton(rd1)
	a.radioLogin = NewColleagueRadioButton(rd2)
	a.textUsername = NewColleagueTextField(edit1)
	a.textPassword = NewColleagueTextField(edit2)
	a.buttonOk = NewColleagueButton(pb1)
	a.buttonCancel = NewColleagueButton(pb2)

	// The initial value of a radio button
	type RadioButtonValue struct{ Mode int }
	radioButtonValue := &RadioButtonValue{0}

	// Create main window
	var mw *walk.MainWindow
	err := declarative.MainWindow{
		AssignTo: &mw,
		Title:    "Mediator Example",
		MinSize:  declarative.Size{Width: 250, Height: 200},
		Size:     declarative.Size{Width: 250, Height: 200},
		Layout:   declarative.VBox{},
		DataBinder: declarative.DataBinder{
			DataSource: radioButtonValue,
		},
		Children: []declarative.Widget{
			declarative.RadioButtonGroupBox{
				DataMember: "Mode",
				Layout:     declarative.HBox{},
				Buttons: []declarative.RadioButton{
					{
						AssignTo:  &a.radioGuest.radioButton,
						Text:      "Guest",
						OnClicked: a.radioGuest.OnClicked,
						Value:     0,
					},
					{
						AssignTo:  &a.radioLogin.radioButton,
						Text:      "Login",
						OnClicked: a.radioLogin.OnClicked,
						Value:     1,
					},
				},
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{
					declarative.Label{
						Text: "Username:",
					},
					declarative.LineEdit{
						AssignTo:      &a.textUsername.lineEdit,
						MinSize:       declarative.Size{Width: 110, Height: 0},
						OnTextChanged: a.textUsername.OnTextChanged,
					},
				},
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{
					declarative.Label{
						Text: "Password:",
					},
					declarative.LineEdit{
						AssignTo:      &a.textPassword.lineEdit,
						MinSize:       declarative.Size{Width: 110, Height: 0},
						PasswordMode:  true,
						OnTextChanged: a.textPassword.OnTextChanged,
					},
				},
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{
					declarative.PushButton{
						AssignTo:  &a.buttonOk.pushButton,
						Text:      "OK",
						OnClicked: a.buttonOk.OnClicked,
					},
					declarative.PushButton{
						AssignTo:  &a.buttonCancel.pushButton,
						Text:      "Cancel",
						OnClicked: a.buttonCancel.OnClicked,
					},
				},
			},
		},
	}.Create()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Set mediators
	a.radioLogin.SetMediator(a)
	a.radioGuest.SetMediator(a)
	a.textUsername.SetMediator(a)
	a.textPassword.SetMediator(a)
	a.buttonOk.SetMediator(a)
	a.buttonCancel.SetMediator(a)

	// Change colleagues in the initial state (the "Guest" radio button is selected).
	a.ColleagueChanged()

	mw.Run()
	// ˄
}

// Change enable/disable of the Colleagues when notified from the Mediators.
func (a *AppLogin) ColleagueChanged() {
	// ˅
	if a.buttonOk.IsPressed() || a.buttonCancel.IsPressed() {
		os.Exit(0)
	} else {
		if a.radioGuest.IsSelected() { // Guest mode
			a.textUsername.SetActivation(false)
			a.textPassword.SetActivation(false)
			a.buttonOk.SetActivation(true)
		} else { // Login mode
			a.textUsername.SetActivation(true)
			a.textPassword.SetActivation(true)

			// Judge whether the changed Colleage is enabled or disabled
			if !a.textUsername.IsEmpty() && !a.textPassword.IsEmpty() {
				a.buttonOk.SetActivation(true)
			} else {
				a.buttonOk.SetActivation(false)
			}
		}
	}
	// ˄
}

// ˅

// ˄
