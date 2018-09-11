// ˅
package main

// ˄

type BookShelfIterator struct {
	// ˅

	// ˄

	index int

	bookShelf *BookShelf

	// ˅

	// ˄
}

func NewBookShelfIterator(bookShelf *BookShelf) *BookShelfIterator {
	// ˅
	return &BookShelfIterator{0, bookShelf}
	// ˄
}

func (self *BookShelfIterator) HasNext() bool {
	// ˅
	return self.index < self.bookShelf.numberOfBooks
	// ˄
}

func (self *BookShelfIterator) Next() interface{} {
	// ˅
	book := self.bookShelf.books[self.index]
	self.index++
	return book
	// ˄
}

// ˅

// ˄
