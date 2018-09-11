// ˅
package main

// ˄

type NightState struct {
	// ˅

	// ˄

	// ˅

	// ˄
}

func NewNightState() *NightState {
	// ˅
	return &NightState{}
	// ˄
}

// Set time
func (self *NightState) SetTime(context Context, hour int) {
	// ˅
	if 9 <= hour && hour < 17 {
		context.ChangeState(NewDaytimeState())
	}
	// ˄
}

// Use a safe
func (self *NightState) UseSafe(context Context) {
	// ˅
	context.CallSecurityGuardsRoom("Emergency: Use a safe at night!")
	// ˄
}

// Sound a emergency bell
func (self *NightState) SoundBell(context Context) {
	// ˅
	context.CallSecurityGuardsRoom("Sound a emergency bell at night")
	// ˄
}

// Make a normal call
func (self *NightState) Call(context Context) {
	// ˅
	context.RecordSecurityLog("Record a night call")
	// ˄
}

func (self *NightState) ToString() string {
	// ˅
	return "[Night]"
	// ˄
}

// ˅

// ˄
