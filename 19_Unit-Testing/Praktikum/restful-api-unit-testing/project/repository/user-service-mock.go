package repository

import (
	"errors"
	"project/models"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

var data = []models.User{
	{
		ID:       1,
		Name:     "Anna",
		Email:    "anna@gmail.com",
		Password: "pass123",
	},
	{
		ID:       2,
		Name:     "Fildza",
		Email:    "fildza@gmail.com",
		Password: "newpass123"},
}

func (um *UserRepositoryMock) CreateUserController(user *models.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("All fields must be filled")
	} else {
		return nil
	}
}

func (um *UserRepositoryMock) GetUserController(id int) (error, interface{}) {
	found := 0
	for _, val := range data {
		if val.ID == id {
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("user not found"), nil
	}

	return nil, data[id-1]
}

func (um *UserRepositoryMock) GetUsersController() (error, interface{}) {
	if data == nil {
		return errors.New("empty data"), nil
	} else {
		return nil, data
	}
}

func (um *UserRepositoryMock) DeleteUserController(id int) error {
	found := 0
	for _, val := range data {
		if val.ID == id {
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (um *UserRepositoryMock) UpdateUserController(user *models.User, id int) error {

	if user.Name == "" && user.Email == "" && user.Password == "" {
		return errors.New("All fields must be filled")
	}

	found := 0
	for _, val := range data {
		if val.ID == id {
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (um *UserRepositoryMock) LoginUserController(user *models.User) (error, string) {

	if user.Email == "" && user.Password == "" {
		return errors.New("All fields must be filled"), ""
	}

	var index int
	found := 0
	for i, val := range data {
		if val.Email == user.Email {
			index = i
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("user not found"), ""
	}

	if data[index].Password != user.Password {
		return errors.New("login failed"), ""
	}

	*user = data[index]

	return nil, ""
}
