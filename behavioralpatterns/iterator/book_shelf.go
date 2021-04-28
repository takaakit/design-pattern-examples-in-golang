// ˅
package main

// ˄

type BookShelf struct {
	// ˅

	// ˄

	numberOfBooks int

	books []*Book

	// ˅

	// ˄
}

func NewBookShelf(maxsize int) *BookShelf {
	// ˅
	bookShelf := &BookShelf{}
	bookShelf.numberOfBooks = 0
	bookShelf.books = make([]*Book, maxsize)
	return bookShelf
	// ˄
}

func (self *BookShelf) Iterator() Iterator {
	// ˅
	return NewBookShelfIterator(self)
	// ˄
}

func (self *BookShelf) GetAt(index int) *Book {
	// ˅
	return self.books[index]
	// ˄
}

func (self *BookShelf) Add(book *Book) {
	// ˅
	self.books[self.numberOfBooks] = book
	self.numberOfBooks++
	// ˄
}

// ˅

// ˄
