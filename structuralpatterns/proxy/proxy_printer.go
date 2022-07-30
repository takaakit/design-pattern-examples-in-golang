// ˅
package proxy

// ˄

// ProxyPrinter forwards requests to RealPrinter when appropriate.
type ProxyPrinter struct {
	// ˅

	// ˄

	currentName string

	// A printer that actually prints
	real *RealPrinter

	// ˅

	// ˄
}

func NewProxyPrinter(name string) *ProxyPrinter {
	// ˅
	return &ProxyPrinter{currentName: name, real: nil}
	// ˄
}

func (p *ProxyPrinter) GetName() string {
	// ˅
	if p.real != nil {
		return p.real.GetName()
	} else {
		return p.currentName
	}
	// ˄
}

func (p *ProxyPrinter) ChangeName(name string) {
	// ˅
	if p.real != nil {
		p.real.ChangeName(name)
	}

	p.currentName = name
	// ˄
}

func (p *ProxyPrinter) Output(content string) {
	// ˅
	// Check to see if the the RealPrinter had been created, create it if necessary.
	if p.real == nil {
		p.real = NewRealPrinter(p.currentName)
	}

	p.real.Output(content)
	// ˄
}

// ˅

// ˄
