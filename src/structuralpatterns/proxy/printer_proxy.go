// ˅
package main

// ˄

type PrinterProxy struct {
	// ˅

	// ˄

	currentName string

	// A printer that actually prints
	real *RealPrinter

	// ˅

	// ˄
}

func NewPrinterProxy(name string) *PrinterProxy {
	// ˅
	printerProxy := &PrinterProxy{}
	printerProxy.currentName = name
	printerProxy.real = nil
	return printerProxy
	// ˄
}

func (self *PrinterProxy) GetPrinterName() string {
	// ˅
	return self.currentName
	// ˄
}

func (self *PrinterProxy) SetPrinterName(value string) {
	// ˅
	if self.real != nil {
		self.real.printerName = value
	}
	self.currentName = value
	// ˄
}

func (self *PrinterProxy) Output(content string) {
	// ˅
	self.createPrinter()
	if self.real != nil {
		self.real.Output(content)
	}
	// ˄
}

// Create an actual printer
func (self *PrinterProxy) createPrinter() {
	// ˅
	if self.real == nil {
		self.real = NewRealPrinter(self.currentName)
	}
	// ˄
}

// ˅

// ˄
