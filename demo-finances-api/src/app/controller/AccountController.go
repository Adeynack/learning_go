package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/adeynack/learning_go/demo-finances-api/src/app/service"
	"github.com/adeynack/learning_go/demo-finances-api/src/app/model"
	"net/http"
)

type AccountController struct {
	GetAccountList gin.HandlerFunc
	GetAccountById gin.HandlerFunc
}

func NewAccountController(bs *service.BookService ,as *service.AccountService) *AccountController {
	return &AccountController{
		GetAccountList: bs.WithBook(getAccountList),
		GetAccountById: as.WithAccount(getAccountById),
	}
}

func getAccountList(c *gin.Context, book *model.Book) {
	c.String(http.StatusOK, "List of accounts for book %v", book.Id)
}

func getAccountById(c *gin.Context, book *model.Book, accountId int64) {
	c.String(http.StatusOK, "Account %v in book %v", accountId, book.Id)
}
