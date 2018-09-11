// ˅
package main

// ˄

type Director struct {
	// ˅

	// ˄

	builder Builder

	// ˅

	// ˄
}

func NewDirector(builder Builder) *Director {
	// ˅
	return &Director{builder}
	// ˄
}

// Construct a document
func (self *Director) Build() {
	// ˅
	self.builder.CreateTitle("Greeting")                                           // Title
	self.builder.CreateSection("Morning and Afternoon")                            // Section
	self.builder.CreateItems([]string{"Good morning.", "Hello."})                  // Items
	self.builder.CreateSection("Evening")                                          // Other section
	self.builder.CreateItems([]string{"Good evening.", "Good night.", "Goodbye."}) // Other items
	self.builder.Close()
	// ˄
}

// ˅

// ˄
