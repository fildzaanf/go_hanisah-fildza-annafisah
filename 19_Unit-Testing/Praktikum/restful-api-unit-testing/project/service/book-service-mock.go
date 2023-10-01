package service

import (
	"errors"
	"project/models"

	"github.com/stretchr/testify/mock"
)

type BookRepositoryMock struct {
	Mock mock.Mock
}

var dataBook = []models.Book{
	{
		ID:        1,
		Title:     "Title 1",
		Author:    "Author 1",
		Publisher: "Publisher 1",
	},
	{
		ID:        2,
		Title:     "Title 2",
		Author:    "Author 2",
		Publisher: "Publisher 2",
	},
}

func (um *BookRepositoryMock) CreateBookController(book *models.Book) error {

	if book.Title == "" || book.Author == "" || book.Publisher == "" {
		return errors.New("All fields must be filled")
	} else {
		return nil
	}
}

func (um *BookRepositoryMock) GetBookController(id int) (error, interface{}) {

	found := 0
	for _, val := range dataBook {
		if val.ID == id {
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("Book not found"), nil
	}

	return nil, dataBook[id-1]
}

func (um *BookRepositoryMock) GetBooksController() (error, interface{}) {

	if dataBook == nil {
		return errors.New("empty data"), nil
	} else {
		return nil, dataBook
	}
}

func (um *BookRepositoryMock) DeleteBookController(id int) error {

	found := 0
	for _, val := range dataBook {
		if val.ID == id {
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("book not found")
	}

	return nil
}

func (um *BookRepositoryMock) UpdateBookController(book *models.Book, id int) error {

	if book.Author == "" && book.Publisher == "" && book.Title == "" {
		return errors.New("All fields must be filled")
	}

	found := 0
	for _, val := range dataBook {
		if val.ID == id {
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("book not found")
	}

	return nil
}
