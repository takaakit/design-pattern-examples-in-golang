// ˅
package main

import "bytes"

// ˄

type PlainTextBuilder struct {
	// ˅

	// ˄

	// String to output
	result string

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
func (self *PlainTextBuilder) CreateTitle(title string) {
	// ˅
	self.buffer.WriteString("--------------------------------\n") // Decoration line
	self.buffer.WriteString("[" + title + "]\n")                  // Title
	self.buffer.WriteString("\n")                                 // Blank line
	// ˄
}

// Make a section of plain text
func (self *PlainTextBuilder) CreateSection(section string) {
	// ˅
	self.buffer.WriteString("* " + section + "\n") // Section
	self.buffer.WriteString("\n")                  // Blank line
	// ˄
}

// Make items of plain text
func (self *PlainTextBuilder) CreateItems(items []string) {
	// ˅
	for i := 0; i < len(items); i++ {
		self.buffer.WriteString("  - " + items[i] + "\n") // Items
	}
	self.buffer.WriteString("\n") // Blank line
	// ˄
}

func (self *PlainTextBuilder) Close() {
	// ˅
	self.buffer.WriteString("--------------------------------\n") // Decoration line
	self.result = self.buffer.String()
	// ˄
}

// ˅

// ˄
