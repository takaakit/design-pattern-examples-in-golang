// ˅
package main

// ˄

// Command to paint a single point
type PaintingCommand struct {
	// ˅

	// ˄

	// Painting position x
	paintingPosX int

	// Painting position y
	paintingPosY int

	paintingTarget PaintingTarget

	// ˅

	// ˄
}

func NewPaintingCommand(paintingTarget PaintingTarget, paintingPosX int, paintingPosY int) *PaintingCommand {
	// ˅
	paintingCommand := &PaintingCommand{}
	paintingCommand.paintingTarget = paintingTarget
	paintingCommand.paintingPosX = paintingPosX
	paintingCommand.paintingPosY = paintingPosY
	return paintingCommand
	// ˄
}

func (self *PaintingCommand) Execute() {
	// ˅
	self.paintingTarget.Paint(self.paintingPosX, self.paintingPosY)
	// ˄
}

// ˅

// ˄
