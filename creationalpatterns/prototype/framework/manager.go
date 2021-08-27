// ˅
package prototype

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
	return &Manager{displays: map[string]Display{}}
	// ˄
}

func (m *Manager) RegisterDisplay(displayName string, display Display) {
	// ˅
	m.displays[displayName] = display
	// ˄
}

func (m *Manager) GetDisplay(displayName string) Display {
	// ˅
	display := m.displays[displayName]
	return display.Clone()
	// ˄
}

// ˅

// ˄
