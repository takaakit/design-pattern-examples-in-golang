// ˅
package factory

import (
	"fmt"
	"os"
	"path/filepath"
)

// ˄

type Page struct {
	// ˅

	// ˄

	IPage

	Title string

	Author string

	Contents []Item

	// ˅

	// ˄
}

func NewPage(title string, author string) *Page {
	// ˅
	return &Page{Title: title, Author: author, Contents: []Item{}}
	// ˄
}

func (p *Page) Add(item Item) {
	// ˅
	p.Contents = append(p.Contents, item)
	// ˄
}

// Client-Specified Self pattern
func (p *Page) Output(iPage IPage) {
	// ˅
	fileName := p.Title + ".html"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	file.Write(([]byte)(iPage.ToHTML()))
	fmt.Println(fileName + " has been created.")
	currentDir, _ := os.Getwd()
	fmt.Println("Output File: " + filepath.Join(currentDir, fileName))
	// ˄
}

// ˅

// ˄
