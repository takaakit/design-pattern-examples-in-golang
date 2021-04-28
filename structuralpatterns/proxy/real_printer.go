// ˅
package main

import (
	"fmt"
	"time"
)

// ˄

type RealPrinter struct {
	// ˅

	// ˄

	printerName string

	// ˅

	// ˄
}

func NewRealPrinter(name string) *RealPrinter {
	// ˅
	realPrinter := &RealPrinter{}
	realPrinter.printerName = name

	realPrinter.heavyTask("Creating an instance(" + realPrinter.printerName + ") of the Printer")

	return realPrinter
	// ˄
}

// Display a content with the name
func (self *RealPrinter) Output(content string) {
	// ˅
	fmt.Println("=== " + self.printerName + " ===")
	fmt.Println(content)
	// ˄
}

// Heavy task (Please think so...)
func (self *RealPrinter) heavyTask(message string) {
	// ˅
	fmt.Print(message)
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println("Done.")
	// ˄
}

// ˅

// ˄
