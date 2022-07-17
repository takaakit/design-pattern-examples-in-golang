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
	for i := 0; i < len(h.pastCommands); i++ {
		h.pastCommands[i].Execute()
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
	var tmpCommands []Command
	for i := 0; i < len(h.pastCommands)-1; i++ {
		tmpCommands = append(tmpCommands, h.pastCommands[i])
	}
	h.pastCommands = tmpCommands
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
