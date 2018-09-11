package main

import (
	"fmt"
)

// Add books in the bookshelf and display the names of the books in turn.

func main() {
	bookShelf := NewBookShelf(5)
	bookShelf.Add(NewBook("Design Patterns: Elements of Reusable Object-Oriented Software"))
	bookShelf.Add(NewBook("The Object Primer: Agile Model-Driven Development with UML 2.0"))
	bookShelf.Add(NewBook("Software Systems Architecture: Working With Stakeholders Using Viewpoints and Perspectives"))
	bookShelf.Add(NewBook("A Practical Guide to SysML: The Systems Modeling Language"))
	bookShelf.Add(NewBook("A Pattern Language: Towns, Buildings, Construction"))

	var it = bookShelf.Iterator()
	for it.HasNext() {
		var book = it.Next()
		fmt.Println(book.(*Book).title)
	}
}
