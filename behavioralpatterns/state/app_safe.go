// ˅
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// ˄

type AppSafe struct {
	// ˅

	// ˄

	// Current time
	textTime *walk.LineEdit

	// Display of security center
	textMessage *walk.TextEdit

	// Current state
	state State

	// ˅

	// ˄
}

func NewAppSafe() *AppSafe {
	// ˅
	var le1 *walk.LineEdit
	var te1 *walk.TextEdit
	var pb1 *walk.PushButton
	var pb2 *walk.PushButton
	var pb3 *walk.PushButton
	var pb4 *walk.PushButton

	appSafe := &AppSafe{}
	appSafe.textTime = le1
	appSafe.textMessage = te1
	appSafe.state = NewDaytimeState()

	MW := MainWindow{
		Title:   "State Example",
		MinSize: Size{300, 300},
		Layout:  VBox{},
		Children: []Widget{
			LineEdit{
				AssignTo: &appSafe.textTime,
				ReadOnly: true,
			},
			TextEdit{
				AssignTo: &appSafe.textMessage,
				ReadOnly: true,
				VScroll:  true,
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						AssignTo:  &pb1,
						Text:      "Use a safe",
						OnClicked: appSafe.useSafe,
					},
					PushButton{
						AssignTo:  &pb2,
						Text:      "Sound an emergency bell",
						OnClicked: appSafe.soundBell,
					},
					PushButton{
						AssignTo:  &pb3,
						Text:      "Make a call",
						OnClicked: appSafe.call,
					},
					PushButton{
						AssignTo:  &pb4,
						Text:      "Exit",
						OnClicked: appSafe.exit,
					},
				},
			},
		},
	}

	go appSafe.countTime()

	if _, err := MW.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return appSafe
	// ˄
}

// Set time
func (self *AppSafe) SetTime(hour int) {
	// ˅
	var clockString string = "Current Time : "
	if hour < 10 {
		clockString += "0" + strconv.Itoa(hour) + ":00"
	} else {
		clockString += strconv.Itoa(hour) + ":00"
	}
	fmt.Println(clockString)
	if self.textTime != nil {
		self.textTime.SetText(clockString)
	}
	self.state.SetTime(self, hour)
	// ˄
}

// Change state
func (self *AppSafe) ChangeState(state State) {
	// ˅
	fmt.Println("The state changed from " + self.state.ToString() + " to " + state.ToString())
	self.state = state
	// ˄
}

// Call a security guard room
func (self *AppSafe) CallSecurityGuardsRoom(msg string) {
	// ˅
	self.textMessage.AppendText("call! " + msg + "\r\n")
	// ˄
}

// Record security log
func (self *AppSafe) RecordSecurityLog(msg string) {
	// ˅
	self.textMessage.AppendText("record ... " + msg + "\r\n")
	// ˄
}

func (self *AppSafe) useSafe() {
	// ˅
	self.state.UseSafe(self)
	// ˄
}

func (self *AppSafe) soundBell() {
	// ˅
	self.state.SoundBell(self)
	// ˄
}

func (self *AppSafe) call() {
	// ˅
	self.state.Call(self)
	// ˄
}

func (self *AppSafe) exit() {
	// ˅
	os.Exit(0)
	// ˄
}

func (self *AppSafe) countTime() {
	// ˅
	for {
		for hour := 0; hour < 24; hour++ {
			self.SetTime(hour) // Set the time
			time.Sleep(1 * time.Second)
		}
	}
	// ˄
}

// ˅

// ˄
