// ˅
package iterator

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
	return &BookShelf{numberOfBooks: 0, books: make([]*Book, maxsize)}
	// ˄
}

func (b *BookShelf) Iterator() Iterator {
	// ˅
	return NewBookShelfIterator(b)
	// ˄
}

func (b *BookShelf) GetAt(index int) *Book {
	// ˅
	return b.books[index]
	// ˄
}

func (b *BookShelf) Add(book *Book) {
	// ˅
	b.books[b.numberOfBooks] = book
	b.numberOfBooks++
	// ˄
}

func (b *BookShelf) NumberOfBooks() int {
	// ˅
	return b.numberOfBooks
	// ˄
}

// ˅

// ˄
