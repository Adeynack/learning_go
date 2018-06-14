package main

import (
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-service/controller"
	"github.com/gin-gonic/gin"
	"net/http"
<<<<<<< HEAD:demo-finances-api/src/finances-service/routes.go
=======
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-service/controller"
>>>>>>> master:demo-finances-api/src/finances-service/routes.go
)

func registerRoutes(
	bookController *BookController,
	accountController *AccountController,
) *gin.Engine {
	r := gin.Default()

	admin := r.Group("/admin")
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	basicAuthAccounts := gin.Accounts{
		"admin": "admin",
	}

	books := r.Group("/books", gin.BasicAuth(basicAuthAccounts))
	books.GET("", bookController.GetBookList())
	books.POST("", bookController.CreateBook())

	bookById := books.Group("/:bookId")
	bookById.GET("", bookController.GetBookById())
	bookById.PUT("", bookController.UpdateBookById())
	bookById.DELETE("", bookController.DeleteBookById())

	accounts := bookById.Group("accounts")
	accounts.GET("", accountController.GetAccountList())

	accountById := accounts.Group("/:accountId")
	accountById.GET("", accountController.GetAccountById())

	return r
}
