// ˅
package main

// ˄

type IFactory interface {
	CreateProduct(owner string) Product

	RegisterProduct(product Product)

	// ˅

	// ˄
}

// ˅

// ˄
