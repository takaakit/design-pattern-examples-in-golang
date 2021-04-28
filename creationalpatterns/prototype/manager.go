// ˅
package main

// ˄

type Manager struct {
	// ˅

	// ˄

	displays map[string]Display

	// ˅

	// ˄
}

func NewManager() *Manager {
	// ˅
	manager := &Manager{}
	manager.displays = map[string]Display{}
	return manager
	// ˄
}

func (self *Manager) RegisterDisplay(displayName string, display Display) {
	// ˅
	self.displays[displayName] = display
	// ˄
}

func (self *Manager) GetDisplay(displayName string) Display {
	// ˅
	display := self.displays[displayName]
	return display.CreateClone()
	// ˄
}

// ˅

// ˄
