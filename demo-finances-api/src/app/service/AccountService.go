package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
	"net/http"
	. "github.com/adeynack/learning_go/demo-finances-api/src/app/model"
)

type AccountService struct {
	// Dependencies
	BookService *BookService
}

func (service AccountService) WithAccount(f func(c *gin.Context, book *Book, accountId int64)) func(*gin.Context) {
	return service.BookService.WithBook(func(c *gin.Context, book *Book) {
		rawAccountId := c.Param("accountId")
		accountId, err := strconv.ParseInt(rawAccountId, 10, 64)
		if err != nil {
			fmt.Printf("Unable to parse accountId `%v`: %v\n", rawAccountId, err)
			c.String(http.StatusNotFound, "No account with ID `%v` exists.", rawAccountId)
			return
		}
		f(c, book, accountId)
	})
}
