package book

import (
	"context"

	"github.com/vier21/go-book-api/pkg/services/book/model"
	"github.com/vier21/go-book-api/utils"
)

type BookRepo interface {
	StoreBook(context.Context, model.Book) utils.Result
	DeleteBook(context.Context, model.Book) utils.Result
	UpdateBook(context.Context, model.Book) utils.Result
}

type BookService interface {
	GetAllBook(context.Context) <-chan utils.Result
	GetBookByID(context.Context, string) <-chan utils.Result
}
