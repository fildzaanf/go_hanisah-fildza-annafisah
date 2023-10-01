package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"project/models"
	"project/service"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// create new book valid
func TestCreateBookControllerValid(t *testing.T) {

	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	var books = models.Book{
		ID:        1,
		Title:     "Title 1",
		Author:    "Author 1",
		Publisher: "Publisher 1",
	}

	bookRepository.Mock.On("CreateBookController", &books).Return(nil)

	e := echo.New()

	requestData, _ := json.Marshal(books)
	request := httptest.NewRequest(http.MethodPost, "/books/", bytes.NewReader(requestData))
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	CreateBookController(c)
	assert.Equal(t, http.StatusOK, response.Code)
}

// create new book invalid
func TestCreateBookControllerInvalid(t *testing.T) {

	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	books := models.Book{
		ID:    1,
		Title: "book",
	}

	bookRepository.Mock.On("CreateBookController", &books).Return(nil)

	e := echo.New()

	requestData, _ := json.Marshal(books)
	request := httptest.NewRequest(http.MethodPost, "/books/", bytes.NewReader(requestData))
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	CreateBookController(c)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

// get book by id valid
func TestGetBookControllerValid(t *testing.T) {

	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	bookRepository.Mock.On("GetBookController", 1).Return(nil)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/books/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	GetBookController(c)
	assert.Equal(t, http.StatusOK, response.Code)
}

// get book by id invalid
func TestGetBookControllerInvalid(t *testing.T) {

	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	bookRepository.Mock.On("GetBookController", 5).Return(nil)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/books/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("5")

	GetBookController(c)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

// get all books valid
func TestGetBooksController(t *testing.T) {

	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	bookRepository.Mock.On("GetBooksController").Return(nil)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/books/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	GetBooksController(c)
	assert.Equal(t, http.StatusOK, response.Code)
}

// get all books invalid
func TestGetBooksControllerInvalid(t *testing.T) {

	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	expectedError := errors.New("failed to retrieve books")
	bookRepository.Mock.On("GetBooksController").Return(expectedError)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/books/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	GetBooksController(c)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

// delete book by id valid
func TestDeleteBookControllerValid(t *testing.T) {

	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	bookRepository.Mock.On("DeleteBookController", 1).Return(nil)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/books/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	DeleteBookController(c)
	assert.Equal(t, http.StatusOK, response.Code)
}

// delete book by id invalid
func TestDeleteBookControllerInvalid(t *testing.T) {

	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	bookRepository.Mock.On("DeleteBookController", 5).Return(nil)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/books/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("5")

	DeleteBookController(c)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

// update book by id valid
func TestUpdateBookControllerValid(t *testing.T) {

	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	books := models.Book{
		Title: "books update",
	}

	bookRepository.Mock.On("UpdateBookController", &books, 1).Return(nil)

	e := echo.New()
	requestData, _ := json.Marshal(books)
	request := httptest.NewRequest(http.MethodPut, "/books/", bytes.NewReader(requestData))
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	UpdateBookController(c)
	assert.Equal(t, http.StatusOK, response.Code)
}

// delete book by id invalid
func TestUpdateBookControllerInvalid(t *testing.T) {

	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	books := models.Book{}

	bookRepository.Mock.On("UpdateBookController", &books, 1).Return(nil)

	e := echo.New()
	requestData, _ := json.Marshal(books)
	request := httptest.NewRequest(http.MethodPut, "/books/", bytes.NewReader(requestData))
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	UpdateBookController(c)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}
