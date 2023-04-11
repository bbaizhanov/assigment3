package funcii

import (
	"Bookstore/objects"
	"fmt"
	"sort"
	"strconv"
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
		Id, err := strconv.Atoi(book.Id)
		if err != nil {
			fmt.Println("vse ne po kaifu")
		}
		if Id == id {
			ans = book
			fmt.Println(book.Title, book.Cost, book.Desc)
		}
	}
	return ans
}

func (bookshelf *Bookshelf) Search(title string) objects.Book {
	var ans objects.Book
	for _, book := range bookshelf.books {
		if book.Title == title {
			ans = book
		}
	}
	return ans
}

func (bookshelf *Bookshelf) GetAll() []objects.Book {
	return bookshelf.books
}

func (bookshelf *Bookshelf) Update(id int, title string, desc string) objects.Book {
	var ans objects.Book
	for _, book := range bookshelf.books {
		Id, err := strconv.Atoi(book.Id)
		if err != nil {
			fmt.Println("Oshibka")
		}
		if Id == id {
			book.Title = title
			book.Desc = desc
			ans = book
		}
	}
	return ans
}

func (bookshelf Bookshelf) Delete(id int) objects.Book {
	var ans objects.Book
	for i, book := range bookshelf.books {
		Id, err := strconv.Atoi(book.Id)
		if err != nil {
			fmt.Println("vse ne po kaifu")
		}
		if Id == id {
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
