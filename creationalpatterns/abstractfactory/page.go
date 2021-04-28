// ˅
package main

import (
	"fmt"
	"os"
)

// ˄

type Page struct {
	// ˅

	// ˄

	title string

	author string

	contents []Item

	// ˅

	// ˄
}

func NewPage(title string, author string) *Page {
	// ˅
	page := &Page{}
	page.title = title
	page.author = author
	return page
	// ˄
}

func (self *Page) Add(item Item) {
	// ˅
	self.contents = append(self.contents, item)
	// ˄
}

func (self *Page) Output(context string) {
	// ˅
	var fileName = self.title + ".html"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	file.Write(([]byte)(context))
	fmt.Println(fileName + " has been created.")
	// ˄
}

// ˅

// ˄
