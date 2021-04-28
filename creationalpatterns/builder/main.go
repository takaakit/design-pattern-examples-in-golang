package main

import (
	"fmt"
	"os"
)

// Create documents in HTML format and text format.

func main() {
	if len(os.Args) != 2 {
		showUsage()
	} else if os.Args[1] == "plain" {
		plainTextBuilder := NewPlainTextBuilder()
		director := NewDirector(plainTextBuilder)
		director.Build()
		content := plainTextBuilder.result
		fmt.Println(content)
	} else if os.Args[1] == "html" {
		htmlBuilder := NewHTMLBuilder()
		director := NewDirector(htmlBuilder)
		director.Build()
		fileName := htmlBuilder.result
		fmt.Println(fileName + " has been created.")
	} else {
		showUsage()
	}
}

func showUsage() {
	fmt.Println("Usage 1: builder.exe plain      <- Create a document in plain text.")
	fmt.Println("Usage 2: builder.exe html       <- Create a document in HTML.")
}
