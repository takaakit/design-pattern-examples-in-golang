// ˅
package command

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
	return &PaintingCommand{
		paintingPosX:   paintingPosX,
		paintingPosY:   paintingPosY,
		paintingTarget: paintingTarget,
	}
	// ˄
}

func (p *PaintingCommand) Execute() {
	// ˅
	p.paintingTarget.Paint(p.paintingPosX, p.paintingPosY)
	// ˄
}

// ˅

// ˄
