// ˅
package proxy

import (
	"fmt"
	"time"
)

// ˄

type RealPrinter struct {
	// ˅

	// ˄

	name string

	// ˅

	// ˄
}

func NewRealPrinter(name string) *RealPrinter {
	// ˅
	realPrinter := &RealPrinter{name: name}

	realPrinter.heavyTask("Creating an instance(" + realPrinter.name + ") of the Printer")

	return realPrinter
	// ˄
}

func (r *RealPrinter) GetName() string {
	// ˅
	return r.name
	// ˄
}

func (r *RealPrinter) ChangeName(name string) {
	// ˅
	r.name = name
	// ˄
}

// Display a content with the name
func (r *RealPrinter) Output(content string) {
	// ˅
	fmt.Println("==========")
	fmt.Println(content)
	fmt.Println("Printed by " + r.name)
	fmt.Println("==========")
	// ˄
}

// Heavy task (Please think so...)
func (r *RealPrinter) heavyTask(message string) {
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
