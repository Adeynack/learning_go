package main

import (
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-service/controller"
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-service/service"
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-service/repository"
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
