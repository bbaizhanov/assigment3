package controller

import (
	"Bookstore/funcii"
	"Bookstore/objects"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	Find(id int) objects.Book
	Search(title string) objects.Book
	GetAll() []objects.Book
	Update(id int) objects.Book
	Delete(id int) objects.Book
	Add(ctx *gin.Context) objects.Book
	GetSortedAscending() []objects.Book
	GetSortedDescending() []objects.Book
}

type Controller struct {
	bookshelf funcii.Bookshelf
}

func New(bookShelf funcii.Bookshelf) Controller {
	return Controller{
		bookShelf,
	}
}

func (c *Controller) Update(ctx *gin.Context) objects.Book {
	var book objects.Book
	err := ctx.BindJSON(&book)
	if err != nil {
		return objects.Book{}
	}
	return c.bookshelf.Update(book)
}

func (c *Controller) Add(ctx *gin.Context) objects.Book {
	var book objects.Book
	err := ctx.BindJSON(&book)
	if err != nil {
		return objects.Book{}
	}
	book = c.bookshelf.Add(book)
	return book
}

func (c *Controller) Find(ctx *gin.Context) objects.Book {
	var ID int
	err := ctx.BindJSON(&ID)
	if err != nil {
		return objects.Book{}
	}
	return c.bookshelf.Find(ID)
}

func (c *Controller) Search(ctx *gin.Context) objects.Book {
	var Title string
	err := ctx.BindJSON(&Title)
	if err != nil {
		return objects.Book{}
	}
	return c.bookshelf.Search(Title)
}

func (c *Controller) GetAll() []objects.Book {
	return c.bookshelf.GetAll()
}

func (c *Controller) Delete(ctx *gin.Context) objects.Book {
	var ID int
	err := ctx.BindJSON(&ID)
	if err != nil {
		return objects.Book{}
	}
	return c.bookshelf.Delete(ID)
}

func (c *Controller) GetSortedAscending() []objects.Book {
	return c.bookshelf.GetSortedAscending()
}

func (c *Controller) GetSortedDescending() []objects.Book {
	return c.bookshelf.GetSortedDescending()
}
