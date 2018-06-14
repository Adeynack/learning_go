package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
	"net/http"
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-service/repository"
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-service/model"
)

type BookService struct {
	// Dependencies
	BookRepository BookRepository
}

func (service *BookService) WithBook(f func(c *gin.Context, book *Book)) func(*gin.Context) {
	return func(c *gin.Context) {
		rawBookId := c.Param("bookId")
		bookId, err := strconv.ParseInt(rawBookId, 10, 64)
		if err != nil {
			fmt.Printf("Unable to parse bookId `%v`: %v\n", rawBookId, err)
			c.String(http.StatusNotFound, "No book with ID `%v` exists.", rawBookId)
			return
		}
		book := service.BookRepository.GetBookById(bookId)
		if book == nil {
			fmt.Printf("No book with ID `%v` exists.\n", bookId)
			c.String(http.StatusNotFound, "No book with ID `%v` exists.", bookId)
			return
		}
		f(c, book)
	}
}
