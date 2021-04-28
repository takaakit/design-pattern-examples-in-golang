// ˅
package main

import "fmt"

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

func (self *PageCreator) CreateSimpleHomepage(mailAddress string, htmlFileName string) {
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
	// ˄
}

// ˅
var instancePageCreator *PageCreator = nil

// ˄
