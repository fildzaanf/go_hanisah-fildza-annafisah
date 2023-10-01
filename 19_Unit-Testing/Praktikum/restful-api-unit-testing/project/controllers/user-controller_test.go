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

// create new user valid
func TestCreateUserControllerValid(t *testing.T) {

	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	user := models.User{
		ID:       1,
		Name:     "Anna",
		Email:    "anna@gmail.com",
		Password: "pass123",
	}

	userRepository.Mock.On("CreateUserController", &user).Return(nil)

	e := echo.New()

	requestData, _ := json.Marshal(user)
	request := httptest.NewRequest(http.MethodPost, "/users/", bytes.NewReader(requestData))
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	CreateUserController(c)

	assert.Equal(t, http.StatusOK, response.Code)
}

// create new user invalid
func TestCreateUserControllerInvalid(t *testing.T) {

	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	user := models.User{
		ID:   1,
		Name: "Anna",
	}

	userRepository.Mock.On("CreateUserController", &user).Return(nil)

	e := echo.New()

	requestData, _ := json.Marshal(user)
	request := httptest.NewRequest(http.MethodPost, "/users/", bytes.NewReader(requestData))
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	CreateUserController(c)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

// get all users valid
func TestGetUsersControllerValid(t *testing.T) {

	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("GetUsersController").Return(nil)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/users/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	GetUsersController(c)

	assert.Equal(t, http.StatusOK, response.Code)
}

// get all users invalid
func TestGetUsersControllerInvalid(t *testing.T) {

	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("GetUsersController").Return([]models.User{})

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/users/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	GetUsersController(c)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

// get user by id valid
func TestGetUserControllerValid(t *testing.T) {

	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("GetUserController", 1).Return(nil)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/users/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	GetUserController(c)

	assert.Equal(t, http.StatusOK, response.Code)
}

// get user by id invalid
func TestGetUserControllerInvalid(t *testing.T) {

	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("GetUserController", 5).Return(nil)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/users/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("5")

	GetUserController(c)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

// delete user by id valid
func TestDeleteUserControllerValid(t *testing.T) {

	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("DeleteUserController", 1).Return(nil)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/users/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	DeleteUserController(c)

	assert.Equal(t, http.StatusOK, response.Code)
}

// delete user by id invalid
func TestDeleteUserControllerInvalid(t *testing.T) {

	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("DeleteUserController", 5).Return(nil)

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/users/", nil)
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("5")

	DeleteUserController(c)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

// update user by id valid
func TestUpdateUserControllerValid(t *testing.T) {

	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	user := models.User{
		Name: "Anna",
	}

	userRepository.Mock.On("UpdateUserController", &user, 1).Return(nil)

	e := echo.New()
	requestData, _ := json.Marshal(user)
	request := httptest.NewRequest(http.MethodPut, "/users/", bytes.NewReader(requestData))
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	UpdateUserController(c)

	assert.Equal(t, http.StatusOK, response.Code)
}

// update user by id invalid
func TestUpdateUserControllerInvalid(t *testing.T) {

	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	user := models.User{}

	userRepository.Mock.On("UpdateUserController", &user, 5).Return(nil)

	e := echo.New()
	requestData, _ := json.Marshal(user)
	request := httptest.NewRequest(http.MethodPut, "/users/", bytes.NewReader(requestData))
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("5")

	UpdateUserController(c)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

// login user valid
func TestLoginUserControllerValid(t *testing.T) {

	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	user := models.User{
		Email:    "anna@gmail.com",
		Password: "pass123",
	}

	userRepository.Mock.On("LoginUserController", &user).Return(nil)

	e := echo.New()
	requestData, _ := json.Marshal(user)
	request := httptest.NewRequest(http.MethodPost, "/login/", bytes.NewReader(requestData))
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	LoginUserController(c)

	assert.Equal(t, http.StatusOK, response.Code)
}

// login user invalid
func TestLoginUserControllerinvalid(t *testing.T) {

	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	user := models.User{
		Email:    "anna@gmail.com",
		Password: "xxx",
	}

	userRepository.Mock.On("LoginUserController", &user).Return(nil)

	e := echo.New()
	requestData, _ := json.Marshal(user)
	request := httptest.NewRequest(http.MethodPost, "/login/", bytes.NewReader(requestData))
	request.Header.Set("content-type", "application/json")
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)

	LoginUserController(c)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}
