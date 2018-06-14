package repository

import . "github.com/adeynack/learning_go/demo-finances-api/src/finances-service/model"

type BookRepository struct {
}

func (repo BookRepository) GetBookById(bookId int64) *Book {
	for _, b := range books {
		if b.Id == bookId {
			return &b
		}
	}
	return nil
}

var books = []Book{
	{1, "One", 1001},
	{2, "Two", 1001},
	{3, "Three", 1002},
}
