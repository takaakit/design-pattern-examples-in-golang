// ˅
package state

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
func (n *NightState) SetTime(context Context, hour int) {
	// ˅
	if 9 <= hour && hour < 17 {
		context.ChangeState(NewDaytimeState())
	}
	// ˄
}

// Use a safe
func (n *NightState) Use(context Context) {
	// ˅
	context.CallSecurityGuardsRoom("Emergency: Use a safe at night!")
	// ˄
}

// Sound a emergency bell
func (n *NightState) Alarm(context Context) {
	// ˅
	context.CallSecurityGuardsRoom("Sound a emergency bell at night")
	// ˄
}

// Make a normal call
func (n *NightState) Phone(context Context) {
	// ˅
	context.RecordSecurityLog("Record a night call")
	// ˄
}

func (n *NightState) String() string {
	// ˅
	return "[Night]"
	// ˄
}

// ˅

// ˄
