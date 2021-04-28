// ˅
package main

// ˄

type State interface {
	// Set time
	SetTime(context Context, hour int)

	// Use a safe
	UseSafe(context Context)

	// Sound a emergency bell
	SoundBell(context Context)

	// Make a normal call
	Call(context Context)

	ToString() string

	// ˅

	// ˄
}

// ˅

// ˄
