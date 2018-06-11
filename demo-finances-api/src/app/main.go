package main

import (
	"github.com/adeynack/learning_go/demo-finances-api/src/app/controller"
	"github.com/adeynack/learning_go/demo-finances-api/src/app/service"
	"github.com/adeynack/learning_go/demo-finances-api/src/app/repository"
)

func main() {

	bookRepository := repository.BookRepository{}
	bookService := &service.BookService{BookRepository: bookRepository}
	bookController := controller.NewBookController(bookService)
	accountService := &service.AccountService{BookService: bookService}
	accountController := controller.NewAccountController(bookService, accountService)

	r := registerRoutes(bookController, accountController)
	r.Run(":3000")
}
