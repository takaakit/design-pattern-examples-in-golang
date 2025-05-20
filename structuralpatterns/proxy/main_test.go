package proxy

import (
	"fmt"
	"testing"
)

/*
Print on a named printer. Setting and changing the printer name is done by Proxy (ProxyPrinter).
At the time of printing, create an instance of the RealSubject (RealPrinter) for the first time.
*/

func Test(t *testing.T) {
	p := NewProxyPrinter("PRINTER-A")
	fmt.Println("The current printer is " + p.GetName() + ".")
	p.ChangeName("PRINTER-B")
	fmt.Println("The current printer is " + p.GetName() + ".")

	fmt.Println("Print start.")
	p.Output("Nice to meet you.")
	fmt.Println("Print exit.")
}
