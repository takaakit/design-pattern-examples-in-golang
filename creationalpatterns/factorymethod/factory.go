// ˅
package main

// ˄

type Factory struct {
	// ˅

	// ˄

	// ˅

	// ˄
}

func NewFactory() *Factory {
	// ˅
	return &Factory{}
	// ˄
}

func (self *Factory) Create(ifactory IFactory, owner string) Product {
	// ˅
	product := ifactory.CreateProduct(owner)
	ifactory.RegisterProduct(product)
	return product
	// ˄
}

// ˅

// ˄
