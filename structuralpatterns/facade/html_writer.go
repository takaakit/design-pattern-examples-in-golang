// ˅
package facade

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
func (h *HtmlWriter) Heading(title string) {
	// ˅
	h.writer.WriteString("<html>\n")
	h.writer.WriteString("<head><title>" + title + "</title></head>\n")
	h.writer.WriteString("<body>\n")
	h.writer.WriteString("<h1>" + title + "</h1>\n")
	// ˄
}

// Write a paragraph
func (h *HtmlWriter) Paragraph(message string) {
	// ˅
	h.writer.WriteString("<p>" + message + "</p>\n")
	// ˄
}

// Write a link
func (h *HtmlWriter) Anchor(url string, text string) {
	// ˅
	h.writer.WriteString("<a href=\"" + url + "\">" + text + "</a>\n")
	// ˄
}

// Write a mail address
func (h *HtmlWriter) Mailto(mailAddress string, userName string) {
	// ˅
	h.Anchor("mailto:"+mailAddress, userName)
	// ˄
}

func (h *HtmlWriter) Close() {
	// ˅
	h.writer.WriteString("</body>\n")
	h.writer.WriteString("</html>\n")
	h.writer.Close()
	// ˄
}

// ˅

// ˄
