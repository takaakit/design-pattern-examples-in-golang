package iterator

import (
	"fmt"
	"testing"
)

/*
Add books in a bookshelf and display the names of the book in turn.
*/

func Test(t *testing.T) {
	bookShelf := NewBookShelf(5)
	bookShelf.Add(NewBook("Design Patterns: Elements of Reusable Object-Oriented Software"))
	bookShelf.Add(NewBook("The Object Primer: Agile Model-Driven Development with UML 2.0"))
	bookShelf.Add(NewBook("Software Systems Architecture: Working With Stakeholders Using Viewpoints and Perspectives"))
	bookShelf.Add(NewBook("A Practical Guide to SysML: The Systems Modeling Language"))
	bookShelf.Add(NewBook("A Pattern Language: Towns, Buildings, Construction"))

	it := bookShelf.Iterator()
	for it.HasNext() {
		book := it.Next()
		fmt.Println(book.(*Book).Title)
	}
}
