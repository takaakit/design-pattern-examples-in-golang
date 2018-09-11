// ˅
package main

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
	historyCommand := &HistoryCommand{}
	historyCommand.pastCommands = []Command{}
	return historyCommand
	// ˄
}

func (self *HistoryCommand) Execute() {
	// ˅
	for i := 0; i < len(self.pastCommands); i++ {
		self.pastCommands[i].Execute()
	}
	// ˄
}

func (self *HistoryCommand) Add(cmd Command) {
	// ˅
	self.pastCommands = append(self.pastCommands, cmd)
	// ˄
}

// Delete the last command
func (self *HistoryCommand) Undo() {
	// ˅
	var tmpCommands []Command
	for i := 0; i < len(self.pastCommands)-1; i++ {
		tmpCommands = append(tmpCommands, self.pastCommands[i])
	}
	self.pastCommands = tmpCommands
	// ˄
}

// Delete all past commands.
func (self *HistoryCommand) Clear() {
	// ˅
	self.pastCommands = []Command{}
	// ˄
}

// ˅

// ˄
