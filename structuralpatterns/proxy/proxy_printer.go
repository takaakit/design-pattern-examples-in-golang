// ˅
package proxy

// ˄

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
	return p.currentName
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
	if p.real == nil {
		p.real = NewRealPrinter(p.currentName)
	}
	p.real.Output(content)
	// ˄
}

// ˅

// ˄