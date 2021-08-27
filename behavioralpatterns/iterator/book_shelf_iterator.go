// ˅
package iterator

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
	return &BookShelfIterator{index: 0, bookShelf: bookShelf}
	// ˄
}

func (b *BookShelfIterator) HasNext() bool {
	// ˅
	return b.index < b.bookShelf.NumberOfBooks()
	// ˄
}

func (b *BookShelfIterator) Next() interface{} {
	// ˅
	book := b.bookShelf.books[b.index]
	b.index++
	return book
	// ˄
}

// ˅

// ˄
