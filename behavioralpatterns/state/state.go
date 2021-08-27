// ˅
package state

// ˄

type State interface {
	SetTime(context Context, hour int)

	Use(context Context)

	Alarm(context Context)

	Phone(context Context)

	String() string

	// ˅

	// ˄
}

// ˅

// ˄
