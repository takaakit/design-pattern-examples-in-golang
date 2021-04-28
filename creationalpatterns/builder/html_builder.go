// ˅
package main

import (
	"fmt"
	"os"
)

// ˄

type HTMLBuilder struct {
	// ˅

	// ˄

	// File name to create
	result string

	writer *os.File

	// ˅

	// ˄
}

func NewHTMLBuilder() *HTMLBuilder {
	// ˅
	return &HTMLBuilder{}
	// ˄
}

// Make a title of HTML file
func (self *HTMLBuilder) CreateTitle(title string) {
	// ˅
	self.result = title + ".html" // Set a title as a file name
	var err error
	self.writer, err = os.Create(self.result)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	self.writer.WriteString("<html><head><title>" + title + "</title></head><body>\n")
	self.writer.WriteString("<h1>" + title + "</h1>\n")
	// ˄
}

// Make a section of HTML file
func (self *HTMLBuilder) CreateSection(section string) {
	// ˅
	self.writer.WriteString("<p>" + section + "</p>\n") // Write a section
	// ˄
}

// Make items of HTML file
func (self *HTMLBuilder) CreateItems(items []string) {
	// ˅
	self.writer.WriteString("<ul>\n") // Write items
	for i := 0; i < len(items); i++ {
		self.writer.WriteString("<li>" + items[i] + "</li>\n")
	}
	self.writer.WriteString("</ul>\n")
	// ˄
}

func (self *HTMLBuilder) Close() {
	// ˅
	self.writer.WriteString("</body></html>\n")
	defer self.writer.Close() // Close file
	// ˄
}

// ˅

// ˄
