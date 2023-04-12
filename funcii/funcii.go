package funcii

import (
	"Bookstore/objects"
	"sort"
)

type Methods interface {
	Find(id int) objects.Book
	Search(title string) objects.Book
	GetAll() []objects.Book
	Update(id int) objects.Book
	Delete(id int) objects.Book
	Add(book objects.Book) objects.Book
	GetSortedAscending() []objects.Book
	GetSortedDescending() []objects.Book
}

type Bookshelf struct {
	books []objects.Book
}

func New() Bookshelf {
	return Bookshelf{}
}

func (bookshelf *Bookshelf) Find(id int) objects.Book {
	var ans objects.Book
	for _, book := range bookshelf.books {

		if book.Id == id {
			ans = book
		}
	}
	return ans
}

func (bookshelf *Bookshelf) Search(title string) objects.Book {
	var ans objects.Book
	for _, book := range bookshelf.books {
		if book.Title == title {
			ans = book
			break
		}
	}
	return ans
}

func (bookshelf *Bookshelf) GetAll() []objects.Book {
	return bookshelf.books
}

func (bookshelf *Bookshelf) Update(Book objects.Book) objects.Book {
	var ans objects.Book
	for _, book := range bookshelf.books {

		if Book.Id == book.Id {
			book = Book
			ans = book
			break
		}
	}
	return ans
}

func (bookshelf Bookshelf) Delete(id int) objects.Book {
	var ans objects.Book
	for i, book := range bookshelf.books {

		if book.Id == id {
			ans = book
			bookshelf.books = append(bookshelf.books[:i], bookshelf.books[i+1:]...)
			break
		}
	}
	return ans
}

func (bookshelf *Bookshelf) Add(book objects.Book) objects.Book {
	bookshelf.books = append(bookshelf.books, book)
	return book
}

func (bookshelf *Bookshelf) GetSortedDescending() []objects.Book {
	sortedBooks := make([]objects.Book, len(bookshelf.books))
	copy(sortedBooks, bookshelf.books)
	sort.Slice(sortedBooks, func(i, j int) bool {
		return sortedBooks[i].Cost > sortedBooks[j].Cost
	})
	return sortedBooks
}

func (bookshelf *Bookshelf) GetSortedAscending() []objects.Book {
	sortedBooks := make([]objects.Book, len(bookshelf.books))
	copy(sortedBooks, bookshelf.books)
	sort.Slice(sortedBooks, func(i, j int) bool {
		return sortedBooks[i].Cost > sortedBooks[j].Cost
	})
	return sortedBooks
}
