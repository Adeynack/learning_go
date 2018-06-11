package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/adeynack/learning_go/demo-finances-api/src/app/service"
	"github.com/adeynack/learning_go/demo-finances-api/src/app/model"
	"net/http"
)

type BookController struct {
	GetBookList    gin.HandlerFunc
	CreateBook     gin.HandlerFunc
	GetBookById    gin.HandlerFunc
	UpdateBookById gin.HandlerFunc
	DeleteBookById gin.HandlerFunc
}

func NewBookController(s *service.BookService) *BookController {
	return &BookController{
		GetBookList:    getBookList,
		CreateBook:     createBook,
		GetBookById:    s.WithBook(getBookById),
		UpdateBookById: s.WithBook(updateBookById),
		DeleteBookById: s.WithBook(deleteBookById),
	}
}

func getBookList(c *gin.Context) {
	c.String(http.StatusOK, "List of books")
}

func createBook(c *gin.Context) {
	c.String(http.StatusCreated, "New book")
}

func getBookById(c *gin.Context, book *model.Book) {
	c.String(http.StatusOK, "Getting book with ID %v", book.Id)
}

func updateBookById(c *gin.Context, book *model.Book) {
	c.String(http.StatusOK, "Updated book with ID %v", book.Id)
}

func deleteBookById(c *gin.Context, book *model.Book) {
	c.String(http.StatusOK, "Delete book with ID %v", book.Id)
}