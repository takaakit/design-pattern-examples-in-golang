package main

import (
	"fmt"
)

// Visitors visit a file system composed of files and directories, and displays a list of files/directories.

func main() {
	fmt.Println("Create a file system...")
	rootDir := NewDirectory("root")
	homeDir := NewDirectory("home")
	binDir := NewDirectory("bin")
	etcDir := NewDirectory("etc")
	emilyDir := NewDirectory("emily")
	jamesDir := NewDirectory("james")
	oliviaDir := NewDirectory("olivia")

	rootDir.Add(homeDir)
	rootDir.Add(binDir)
	rootDir.Add(etcDir)

	binDir.Add(NewFile("ls", 100))
	binDir.Add(NewFile("mkdir", 50))
	homeDir.Add(emilyDir)
	homeDir.Add(jamesDir)
	homeDir.Add(oliviaDir)

	emilyDir.Add(NewFile("homework.doc", 40))
	jamesDir.Add(NewFile("homework.doc", 50))
	jamesDir.Add(NewFile("app.exe", 60))
	oliviaDir.Add(NewFile("homework.doc", 70))
	oliviaDir.Add(NewFile("app.exe", 80))
	oliviaDir.Add(NewFile("tips.html", 90))

	rootDir.Accept(NewListVisitor())
}
