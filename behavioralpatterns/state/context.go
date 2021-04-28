// ˅
package main

// ˄

type Context interface {
	// Set time
	SetTime(hour int)

	// Change state
	ChangeState(state State)

	// Call a security guard room
	CallSecurityGuardsRoom(msg string)

	// Record security log
	RecordSecurityLog(msg string)

	// ˅

	// ˄
}

// ˅

// ˄
