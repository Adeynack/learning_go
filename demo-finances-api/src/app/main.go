package main

import (
	. "github.com/adeynack/learning_go/demo-finances-api/src/app/controller"
	. "github.com/adeynack/learning_go/demo-finances-api/src/app/service"
	. "github.com/adeynack/learning_go/demo-finances-api/src/app/repository"
)

func main() {

	bookRepository := &BookRepository{}
	bookService := NewBookService(bookRepository)
	bookController := NewBookController(bookService)
	accountService := NewAccountService(bookService)
	accountController := NewAccountController(bookService, accountService)

	r := registerRoutes(bookController, accountController)
	r.Run(":3000")
}
