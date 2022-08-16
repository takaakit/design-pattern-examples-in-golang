// ˅
package command

// ˄

// Holder of the past commands
type HistoryCommand struct {
	// ˅

	// ˄

	// A set of past commands
	pastCommands []Command

	// ˅

	// ˄
}

func NewHistoryCommand() *HistoryCommand {
	// ˅
	return &HistoryCommand{pastCommands: []Command{}}
	// ˄
}

// Execute all past commands
func (h *HistoryCommand) Execute() {
	// ˅
	for _, command := range h.pastCommands {
		command.Execute()
	}
	// ˄
}

func (h *HistoryCommand) Add(cmd Command) {
	// ˅
	h.pastCommands = append(h.pastCommands, cmd)
	// ˄
}

// Delete the last command
func (h *HistoryCommand) Undo() {
	// ˅
	if len(h.pastCommands) >= 1 {
		h.pastCommands = h.pastCommands[:len(h.pastCommands)-1]
	}
	// ˄
}

// Delete all past commands.
func (h *HistoryCommand) Clear() {
	// ˅
	h.pastCommands = []Command{}
	// ˄
}

// ˅

// ˄
