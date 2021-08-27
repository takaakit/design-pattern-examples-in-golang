// ˅
package factorymethod

// ˄

type Factory struct {
	// ˅

	// ˄

	IFactory

	// ˅

	// ˄
}

func NewFactory() *Factory {
	// ˅
	return &Factory{}
	// ˄
}

// GO does not dispatch back to the embedding object. If we wish the embedding object's methods to be called,
// then we must pass in the embedding object as an extra parameter. This is the client-specified self pattern.
// (Evaluating the GO Programming Language with Design Patterns)
func (f *Factory) Create(ifactory IFactory, owner string) Product {
	// ˅
	// Write pre-creation code here, if any.

	product := ifactory.CreateProduct(owner)

	// Write post-creation code here, if any.

	return product
	// ˄
}

// ˅

// ˄
