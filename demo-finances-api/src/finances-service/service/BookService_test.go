package service

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"fmt"
	. "github.com/adeynack/learning_go/demo-finances-api/src/finances-service/model"
)

func TestWithBookIdNone(t *testing.T) {
	service := NewBookService(fakeBookRepository{})
	f := service.WithBookId(func(c *gin.Context, bookId int64) {
		c.String(200, fmt.Sprintf("%v", bookId))
	})

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{
	// No book ID
	}

	f(c)
	assertStatus(t, c, 404)
	assertBody(t, recorder, "No book with ID `` exists.")
}

func TestWithBookId42(t *testing.T) {
	f := fakeBookService()
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{
		gin.Param{Key: "bookId", Value: "42"},
	}

	f(c)
	assertStatus(t, c, 200)
	assertBody(t, recorder, "42")
}

func TestWithBookIdNonInt(t *testing.T) {
	service := NewBookService(fakeBookRepository{})
	f := service.WithBookId(func(c *gin.Context, bookId int64) {
		c.String(200, fmt.Sprintf("%v", bookId))
	})

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{
		gin.Param{Key: "bookId", Value: "abc"},
	}

	f(c)
	assertStatus(t, c, 404)
	assertBody(t, recorder, "No book with ID `abc` exists.")
}

//
// Helper
//

type fakeBookRepository struct {
}

func (repo fakeBookRepository) GetBookById(bookId int64) *Book {
	return nil
}

func fakeBookService() gin.HandlerFunc {
	service := NewBookService(fakeBookRepository{})
	return service.WithBookId(func(c *gin.Context, bookId int64) {
		c.String(200, fmt.Sprintf("%v", bookId))
	})
}

func assertBody(t *testing.T, r *httptest.ResponseRecorder, expected string) {
	if body := r.Body.String(); body != expected {
		if len(expected) == 0 {
			t.Errorf("Expected empty body, got %v", body)
		} else {
			t.Errorf("Expected body %v, got %v", expected, body)
		}
	}
}

func assertStatus(t *testing.T, c *gin.Context, expected int) {
	if status := c.Writer.Status(); status != expected {
		t.Errorf("Expected status %v, got %v", expected, status)
	}
}
