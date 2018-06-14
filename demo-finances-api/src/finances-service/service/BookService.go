package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
	"net/http"
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-service/repository"
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-service/model"
)

type BookService interface {
	WithBookId(f func(c *gin.Context, bookId int64)) gin.HandlerFunc
	WithBook(f func(c *gin.Context, book *Book)) gin.HandlerFunc
}

func NewBookService(bookRepository BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepository,
	}
}

type bookService struct {
	bookRepository BookRepository
}

func (service *bookService) WithBookId(f func(c *gin.Context, bookId int64)) gin.HandlerFunc {
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

func (service *bookService) WithBook(f func(c *gin.Context, book *Book)) gin.HandlerFunc {
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
