package main

import (
	"fmt"
)

// At the time of printing, create an instance of the printer for the first time.
// In order to spend time creating a printer, call a heavy task when creating a printer instance.

func main() {
	p := NewPrinterProxy("Emily's printer")
	fmt.Println("The current printer is " + p.GetPrinterName() + ".")
	p.SetPrinterName("William's printer")
	fmt.Println("The current printer is " + p.GetPrinterName() + ".")
	p.Output("Nice to meet you.")
}
