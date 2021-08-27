// ˅
package builder

import "bytes"

// ˄

type PlainTextBuilder struct {
	// ˅

	// ˄

	buffer bytes.Buffer

	// ˅

	// ˄
}

func NewPlainTextBuilder() *PlainTextBuilder {
	// ˅
	return &PlainTextBuilder{}
	// ˄
}

// Make a title of plain text
func (p *PlainTextBuilder) CreateTitle(title string) {
	// ˅
	p.buffer.WriteString("--------------------------------\n") // Decoration line
	p.buffer.WriteString("[" + title + "]\n")                  // Title
	p.buffer.WriteString("\n")                                 // Blank line
	// ˄
}

// Make a section of plain text
func (p *PlainTextBuilder) CreateSection(section string) {
	// ˅
	p.buffer.WriteString("* " + section + "\n") // Section
	p.buffer.WriteString("\n")                  // Blank line
	// ˄
}

// Make items of plain text
func (p *PlainTextBuilder) CreateItems(items []string) {
	// ˅
	for i := 0; i < len(items); i++ {
		p.buffer.WriteString("  - " + items[i] + "\n") // Items
	}
	p.buffer.WriteString("\n") // Blank line
	// ˄
}

func (p *PlainTextBuilder) Close() {
	// ˅
	p.buffer.WriteString("--------------------------------\n") // Decoration line
	// ˄
}

// String to output
func (p *PlainTextBuilder) GetContent() string {
	// ˅
	return p.buffer.String()
	// ˄
}

// ˅

// ˄
