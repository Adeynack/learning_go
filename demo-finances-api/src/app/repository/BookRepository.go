package repository

import "github.com/adeynack/learning_go/demo-finances-api/src/app/model"

type BookRepository struct {
	// Dependencies
}

func (repo BookRepository) GetBookById(bookId int64) *model.Book {
	for _, b := range books {
		if b.Id == bookId {
			return &b
		}
	}
	return nil
}

var books = []model.Book{
	{1, "One", 1001},
	{2, "Two", 1001},
	{3, "Three", 1002},
}
