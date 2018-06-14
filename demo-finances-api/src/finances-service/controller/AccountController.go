package controller

import (
	"github.com/gin-gonic/gin"
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-service/service"
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-service/model"
	"net/http"
)

type AccountController struct {
	bookService    *BookService
	accountService *AccountService
}

func NewAccountController(bookService *BookService, accountService *AccountService) *AccountController {
	return &AccountController{
		bookService:    bookService,
		accountService: accountService,
	}
}

func (ctrl AccountController) GetAccountList() gin.HandlerFunc {
	return ctrl.bookService.WithBook(func(c *gin.Context, book *Book) {
		c.String(http.StatusOK, "List of accounts for book %v", book.Id)
	})
}

func (ctrl AccountController) GetAccountById() gin.HandlerFunc {
	return ctrl.accountService.WithAccount(func(c *gin.Context, book *Book, accountId int64) {
		c.String(http.StatusOK, "Account %v in book %v", accountId, book.Id)
	})
}
