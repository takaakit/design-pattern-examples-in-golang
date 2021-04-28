// ˅
package main

import (
	"fmt"
	"os"
)

// ˄

type HtmlWriter struct {
	// ˅

	// ˄

	writer *os.File

	// ˅

	// ˄
}

func NewHtmlWriter(htmlFileName string) *HtmlWriter {
	// ˅
	htmlWriter := &HtmlWriter{}
	var err error
	htmlWriter.writer, err = os.Create(htmlFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return htmlWriter
	// ˄
}

// Write a title
func (self *HtmlWriter) Heading(title string) {
	// ˅
	self.writer.WriteString("<html>\n")
	self.writer.WriteString("<head><title>" + title + "</title></head>\n")
	self.writer.WriteString("<body>\n")
	self.writer.WriteString("<h1>" + title + "</h1>\n")
	// ˄
}

// Write a paragraph
func (self *HtmlWriter) Paragraph(message string) {
	// ˅
	self.writer.WriteString("<p>" + message + "</p>\n")
	// ˄
}

// Write a link
func (self *HtmlWriter) Anchor(url string, text string) {
	// ˅
	self.writer.WriteString("<a href=\"" + url + "\">" + text + "</a>\n")
	// ˄
}

// Write a mail address
func (self *HtmlWriter) Mailto(mailAddress string, userName string) {
	// ˅
	self.Anchor("mailto:"+mailAddress, userName)
	// ˄
}

func (self *HtmlWriter) Close() {
	// ˅
	self.writer.WriteString("</body>\n")
	self.writer.WriteString("</html>\n")
	defer self.writer.Close()
	// ˄
}

// ˅

// ˄
