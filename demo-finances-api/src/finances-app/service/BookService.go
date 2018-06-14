package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
	"net/http"
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-app/repository"
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-app/model"
)

type BookService struct {
	bookRepository *BookRepository
}

func NewBookService(bookRepository *BookRepository) *BookService {
	return &BookService{
		bookRepository: bookRepository,
	}
}

func (service *BookService) WithBookId(f func(c *gin.Context, bookId int64)) gin.HandlerFunc {
	return func(c *gin.Context) {
		rawBookId := c.Param("bookId")
		bookId, err := strconv.ParseInt(rawBookId, 10, 64)
		if err != nil {
			fmt.Printf("Unable to parse bookId `%v`: %v\n", rawBookId, err)
			c.String(http.StatusNotFound, "No book with ID `%v` exists.", rawBookId)
			return
		}
		f(c, bookId)
	}
}

func (service *BookService) WithBook(f func(c *gin.Context, book *Book)) gin.HandlerFunc {
	return service.WithBookId(func(c *gin.Context, bookId int64) {
		book := service.bookRepository.GetBookById(bookId)
		if book == nil {
			fmt.Printf("No book with ID `%v` exists.\n", bookId)
			c.String(http.StatusNotFound, "No book with ID `%v` exists.", bookId)
			return
		}
		f(c, book)
	})
}
