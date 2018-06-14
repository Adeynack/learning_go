package controller

import (
	"github.com/gin-gonic/gin"
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-app/service"
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-app/model"
	"net/http"
)

type BookController struct {
	bookService *BookService
}

func NewBookController(bookService *BookService) *BookController {
	return &BookController{
		bookService: bookService,
	}
}

func (ctrl BookController) GetBookList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "List of books")
	}
}

func (ctrl BookController) CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusCreated, "New book")
	}
}

func (ctrl BookController) GetBookById() gin.HandlerFunc {
	return ctrl.bookService.WithBook(func(c *gin.Context, book *Book) {
		c.String(http.StatusOK, "Getting book with ID %v", book.Id)
	})
}

func (ctrl BookController) UpdateBookById() gin.HandlerFunc {
	return ctrl.bookService.WithBook(func(c *gin.Context, book *Book) {
		c.String(http.StatusOK, "Updated book with ID %v", book.Id)
	})
}

func (ctrl BookController) DeleteBookById() gin.HandlerFunc {
	return ctrl.bookService.WithBook(func(c *gin.Context, book *Book) {
		c.String(http.StatusOK, "Delete book with ID %v", book.Id)
	})
}
