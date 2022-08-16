// ˅
package builder

import (
	"fmt"
	"os"
)

// ˄

type HTMLBuilder struct {
	// ˅

	// ˄

	// File name to create
	fileName string

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
func (h *HTMLBuilder) CreateTitle(title string) {
	// ˅
	h.fileName = title + ".html" // Set a title as a file name
	var err error
	h.writer, err = os.Create(h.fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	h.writer.WriteString("<html><head><title>" + title + "</title></head><body>\n")
	h.writer.WriteString("<h1>" + title + "</h1>\n")
	// ˄
}

// Make a section of HTML file
func (h *HTMLBuilder) CreateSection(section string) {
	// ˅
	h.writer.WriteString("<p>" + section + "</p>\n") // Write a section
	// ˄
}

// Make items of HTML file
func (h *HTMLBuilder) CreateItems(items []string) {
	// ˅
	h.writer.WriteString("<ul>\n") // Write items
	for _, item := range items {
		h.writer.WriteString("<li>" + item + "</li>\n")
	}
	h.writer.WriteString("</ul>\n")
	// ˄
}

func (h *HTMLBuilder) Close() {
	// ˅
	h.writer.WriteString("</body></html>\n")
	defer h.writer.Close() // Close file
	// ˄
}

func (h *HTMLBuilder) GetFileName() string {
	// ˅
	return h.fileName
	// ˄
}

// ˅

// ˄
