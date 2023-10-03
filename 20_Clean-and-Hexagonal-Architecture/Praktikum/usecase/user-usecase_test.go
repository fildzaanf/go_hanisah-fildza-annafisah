package usecase

import (
	"Praktikum/entity"
	"Praktikum/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockUserRepository = &repository.MockUserRepository{Mock: mock.Mock{}}
var mockUserUseCase = userUseCase{userRepository: mockUserRepository}

func TestCreateUserController(t *testing.T) {
	data := entity.User{
		ID:       1,
		Name:     "Anna",
		Email:    "anna@gmail.com",
		Password: "pass123",
	}

	mockUserRepository.Mock.On("CreateUserController", &data).Return(data)

	err := mockUserUseCase.CreateUserController(&data)
	assert.Nil(t, err)

}

func TestCreateUserControllerInvalid(t *testing.T) {
	data := entity.User{}

	mockUserRepository.Mock.On("CreateUserController", &data).Return(errors.New("All fields must be filled"))

	err := mockUserUseCase.CreateUserController(&data)
	assert.NotNil(t, err)
}

func TestGetUsersController(t *testing.T) {

	mockUserRepository.Mock.On("GetUsersController").Return()

	err, res := mockUserUseCase.GetUsersController()
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestLoginUserController(t *testing.T) {
	data := entity.User{
		Email:    "anna@gmail.com",
		Password: "pass123",
	}

	mockUserRepository.Mock.On("LoginUserController", &data).Return(nil)

	err, token := mockUserUseCase.LoginUserController(&data)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestLoginUserControllerInvalid(t *testing.T) {
	data := entity.User{
		Email:    "anna@gmail.com",
		Password: "",
	}

	mockUserRepository.Mock.On("LoginUserController", &data).Return(errors.New("user not found"))

	err, token := mockUserUseCase.LoginUserController(&data)
	assert.Equal(t, "", token)
	assert.NotNil(t, err)
}

func TestNewUserUseCase(t *testing.T) {

	dummyRepo := &repository.MockUserRepository{}

	userUseCase := NewUserUseCase(dummyRepo)

	if userUseCase.userRepository != dummyRepo {
		t.Errorf("Expected userUseCase.userRepository to be equal to dummyRepo, but got %v", userUseCase.userRepository)
	}
}
