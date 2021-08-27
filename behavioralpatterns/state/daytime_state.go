// ˅
package state

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
func (d *DaytimeState) SetTime(context Context, hour int) {
	// ˅
	if hour < 9 || 17 <= hour {
		context.ChangeState(NewNightState())
	}
	// ˄
}

// Use a safe
func (d *DaytimeState) Use(context Context) {
	// ˅
	context.RecordSecurityLog("Use a safe in the daytime")
	// ˄
}

// Sound a emergency bell
func (d *DaytimeState) Alarm(context Context) {
	// ˅
	context.CallSecurityGuardsRoom("Sound a emergency bell in the daytime")
	// ˄
}

// Make a normal call
func (d *DaytimeState) Phone(context Context) {
	// ˅
	context.CallSecurityGuardsRoom("Make a normal call in the daytime")
	// ˄
}

func (d *DaytimeState) String() string {
	// ˅
	return "[Daytime]"
	// ˄
}

// ˅

// ˄
