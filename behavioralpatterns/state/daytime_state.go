// ˅
package main

// ˄

type DaytimeState struct {
	// ˅

	// ˄

	// ˅

	// ˄
}

func NewDaytimeState() *DaytimeState {
	// ˅
	return &DaytimeState{}
	// ˄
}

// Set time
func (self *DaytimeState) SetTime(context Context, hour int) {
	// ˅
	if hour < 9 || 17 <= hour {
		context.ChangeState(NewNightState())
	}
	// ˄
}

// Use a safe
func (self *DaytimeState) UseSafe(context Context) {
	// ˅
	context.RecordSecurityLog("Use a safe in the daytime")
	// ˄
}

// Sound a emergency bell
func (self *DaytimeState) SoundBell(context Context) {
	// ˅
	context.CallSecurityGuardsRoom("Sound a emergency bell in the daytime")
	// ˄
}

// Make a normal call
func (self *DaytimeState) Call(context Context) {
	// ˅
	context.CallSecurityGuardsRoom("Make a normal call in the daytime")
	// ˄
}

func (self *DaytimeState) ToString() string {
	// ˅
	return "[Daytime]"
	// ˄
}

// ˅

// ˄
