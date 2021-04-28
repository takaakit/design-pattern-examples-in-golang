package main

import (
	"fmt"
	"os"
)

// First, create instances for displaying large size characters, then display large size character string using that instances.

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: flyweight.exe digits")
		fmt.Println("Ex.  : flyweight.exe 1212123")
	} else {
		bs := NewLargeSizeString(os.Args[1])
		bs.Display()
	}
}
