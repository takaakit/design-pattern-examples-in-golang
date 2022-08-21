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

// Note: This is the client-specified self pattern.
func (p *Page) Output(iPage IPage) {
	// ˅
	fileName := p.Title + ".html"
	file, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}
	defer file.Close()

	_, err2 := file.Write(([]byte)(iPage.ToHTML()))
	if err2 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}
	fmt.Println(fileName + " has been created.")
	currentDir, _ := os.Getwd()
	fmt.Println("Output File: " + filepath.Join(currentDir, fileName))
	// ˄
}

// ˅

// ˄
