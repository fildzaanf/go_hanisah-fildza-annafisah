package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project/models"
	"project/service"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBlogControllerValid(t *testing.T) {

	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	var blogs = models.Blog{
		ID:      1,
		UserId:  1,
		Title:   "Title 1",
		Content: "Content 1",
	}

	blogRepository.Mock.On("CreateBlogController", &blogs).Return(nil)

	e := echo.New()

	requestData, _ := json.Marshal(blogs)
	request := httptest.NewRequest(http.MethodPost, "/blogs/", bytes.NewReader(requestData))
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	CreateBlogController(c)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestCreateBlogControllerInvalid(t *testing.T) {

	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	blogs := models.Blog{
		Title: "blogs",
	}
	blogRepository.Mock.On("CreateBlogController", &blogs).Return(nil)

	e := echo.New()

	requestData, _ := json.Marshal(blogs)
	request := httptest.NewRequest(http.MethodPost, "/blogs/", bytes.NewReader(requestData))
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	CreateBlogController(c)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestGetBlogControllerValid(t *testing.T) {

	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	blogRepository.Mock.On("GetBlogController", 1).Return(nil)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/blogs/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	GetBlogController(c)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetBlogControllerInvalid(t *testing.T) {

	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	blogRepository.Mock.On("GetBlogController", 5).Return(nil)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/blogs/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("5")

	GetBlogController(c)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestGetBlogsController(t *testing.T) {

	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	blogRepository.Mock.On("GetBlogsController").Return(nil)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/blogs/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	GetBlogsController(c)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeleteBlogControllerValid(t *testing.T) {

	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	blogRepository.Mock.On("DeleteBlogController", 1).Return(nil)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/blogs/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	DeleteBlogController(c)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeleteBlogControllerInvalid(t *testing.T) {

	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	blogRepository.Mock.On("DeleteBlogController", 5).Return(nil)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/blogs/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("5")

	DeleteBlogController(c)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestUpdateBlogControllerValid(t *testing.T) {

	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	blogs := models.Blog{
		Title: "blogs updated",
	}

	blogRepository.Mock.On("UpdateBlogController", &blogs, 1).Return(nil)

	e := echo.New()
	requestData, _ := json.Marshal(blogs)
	request := httptest.NewRequest(http.MethodPut, "/blogs/", bytes.NewReader(requestData))
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	UpdateBlogController(c)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUpdateBlogControllerInvalid(t *testing.T) {

	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	blogs := models.Blog{}

	blogRepository.Mock.On("UpdateBlogController", &blogs, 1).Return(nil)

	e := echo.New()
	requestData, _ := json.Marshal(blogs)
	request := httptest.NewRequest(http.MethodPut, "/blogs/", bytes.NewReader(requestData))
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	UpdateBlogController(c)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}
