package controller

import (
	"github.com/gin-gonic/gin"
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-service/service"
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-service/model"
	"net/http"
)

type AccountController struct {
	GetAccountList gin.HandlerFunc
	GetAccountById gin.HandlerFunc
}

func NewAccountController(bs *BookService, as *AccountService) *AccountController {
	return &AccountController{
		GetAccountList: bs.WithBook(getAccountList),
		GetAccountById: as.WithAccount(getAccountById),
	}
}

func getAccountList(c *gin.Context, book *Book) {
	c.String(http.StatusOK, "List of accounts for book %v", book.Id)
}

func getAccountById(c *gin.Context, book *Book, accountId int64) {
	c.String(http.StatusOK, "Account %v in book %v", accountId, book.Id)
}
