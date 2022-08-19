// ˅
package state

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

// ˄

// Safe security system that the security status changes with time.
type AppSafe struct {
	// ˅

	// ˄

	// Current state
	state State

	// Current time
	textTime *walk.LineEdit

	// Display of security center
	textMessage *walk.TextEdit

	// ˅

	// ˄
}

func NewAppSafe() *AppSafe {
	// ˅
	var le *walk.LineEdit
	var te *walk.TextEdit
	var pb1 *walk.PushButton
	var pb2 *walk.PushButton
	var pb3 *walk.PushButton

	appSafe := &AppSafe{
		textTime:    le,
		textMessage: te,
		state:       NewDaytimeState(),
	}

	go appSafe.countTime()

	_, err := declarative.MainWindow{
		Title:   "State Example",
		MinSize: declarative.Size{Width: 300, Height: 300},
		Size:    declarative.Size{Width: 300, Height: 300},
		Layout:  declarative.VBox{},
		Children: []declarative.Widget{
			declarative.LineEdit{
				AssignTo: &appSafe.textTime,
				ReadOnly: true,
			},
			declarative.TextEdit{
				AssignTo: &appSafe.textMessage,
				ReadOnly: true,
				VScroll:  true,
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{
					declarative.PushButton{
						AssignTo:  &pb1,
						Text:      "Use",
						OnClicked: appSafe.pressedUseButton,
					},
					declarative.PushButton{
						AssignTo:  &pb2,
						Text:      "Alarm",
						OnClicked: appSafe.pressedAlarmButton,
					},
					declarative.PushButton{
						AssignTo:  &pb3,
						Text:      "Phone",
						OnClicked: appSafe.pressedPhoneButton,
					},
				},
			},
		},
	}.Run()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return appSafe
	// ˄
}

// Set time
func (a *AppSafe) SetTime(hour int) {
	// ˅
	var clockString string = "Current Time : "
	if hour < 10 {
		clockString += "0" + strconv.Itoa(hour) + ":00"
	} else {
		clockString += strconv.Itoa(hour) + ":00"
	}
	fmt.Println(clockString)
	if a.textTime != nil {
		_ = a.textTime.SetText(clockString)
	}
	a.state.SetTime(a, hour)
	// ˄
}

// Change state
func (a *AppSafe) ChangeState(state State) {
	// ˅
	fmt.Println("The state changed from " + a.state.String() + " to " + state.String())
	a.state = state
	// ˄
}

// Call a security guard room
func (a *AppSafe) CallSecurityGuardsRoom(msg string) {
	// ˅
	a.textMessage.AppendText("call! " + msg + "\r\n")
	// ˄
}

// Record security log
func (a *AppSafe) RecordSecurityLog(msg string) {
	// ˅
	a.textMessage.AppendText("record ... " + msg + "\r\n")
	// ˄
}

func (a *AppSafe) pressedUseButton() {
	// ˅
	a.state.Use(a)
	// ˄
}

func (a *AppSafe) pressedAlarmButton() {
	// ˅
	a.state.Alarm(a)
	// ˄
}

func (a *AppSafe) pressedPhoneButton() {
	// ˅
	a.state.Phone(a)
	// ˄
}

func (a *AppSafe) countTime() {
	// ˅
	for {
		// Advance one hour for every second of real time.
		for hour := 0; hour < 24; hour++ {
			a.SetTime(hour) // Set the time
			time.Sleep(1 * time.Second)
		}
	}
	// ˄
}

// ˅

// ˄
