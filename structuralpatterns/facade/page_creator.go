// ˅
package facade

import (
	"fmt"
	"os"
	"path/filepath"
)

// ˄

type PageCreator struct {
	// ˅

	// ˄

	// ˅

	// ˄
}

func NewPageCreator() *PageCreator {
	// ˅
	if instancePageCreator == nil {
		instancePageCreator = &PageCreator{}
	}
	return instancePageCreator
	// ˄
}

func (p *PageCreator) CreateSimpleHomepage(mailAddress string, htmlFileName string) {
	// ˅
	addressBook := NewDataLibrary()
	data := addressBook.GetProperties("addressbook")
	userName := data[mailAddress]
	writer := NewHtmlWriter(htmlFileName)
	writer.Heading(userName + "'s homepage")
	writer.Paragraph("Welcome to " + userName + "'s homepage.")
	writer.Paragraph("Please email me at this address.")
	writer.Mailto(mailAddress, userName)
	writer.Close()
	fmt.Println(htmlFileName + " is created for " + mailAddress + " (" + userName + ")")
	currentDir, _ := os.Getwd()
	fmt.Println("Output File: " + filepath.Join(currentDir, htmlFileName))
	// ˄
}

// ˅
var instancePageCreator *PageCreator = nil

// ˄
