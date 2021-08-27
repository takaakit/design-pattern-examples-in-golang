// ˅
package builder

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
func (d *Director) Build() {
	// ˅
	d.builder.CreateTitle("Greeting")                                           // Title
	d.builder.CreateSection("Morning and Afternoon")                            // Section
	d.builder.CreateItems([]string{"Good morning.", "Hello."})                  // Items
	d.builder.CreateSection("Evening")                                          // Other section
	d.builder.CreateItems([]string{"Good evening.", "Good night.", "Goodbye."}) // Other items
	d.builder.Close()
	// ˄
}

// ˅

// ˄
