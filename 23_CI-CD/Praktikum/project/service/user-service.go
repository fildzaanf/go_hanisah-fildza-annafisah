package service

import (
	"errors"
	"net/http"
	"project/config"
	"project/middlewares"
	"project/models"

	"github.com/labstack/echo/v4"
)

type UserService interface {
	CreateUserController(user *models.User) error
	GetUserController(id int) (error, interface{})
	GetUsersController() (error, interface{})
	DeleteUserController(id int) error
	UpdateUserController(user *models.User, id int) error
	LoginUserController(user *models.User) (error, string)
}

type UserRepository struct {
	Func UserService
}

var userRepository UserService

func init() {
	ur := &UserRepository{}
	ur.Func = ur

	userRepository = ur
}

func GetUserRepository() UserService {
	return userRepository
}

func SetUserRepository(ur UserService) {
	userRepository = ur
}

// create new user
func (u *UserRepository) CreateUserController(user *models.User) error {

	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("All fields must be filled")
	}

	if err := config.DB.Save(&user); err != nil {
		return err.Error
	}

	return nil
}

// get user by id
func (u *UserRepository) GetUserController(id int) (err error, res interface{}) {

	var user models.User

	if result := config.DB.Where("id = ?", id).First(&user); result.Error != nil {
		return errors.New("user not found"), nil
	}

	return nil, user
}

// get all users
func (u *UserRepository) GetUsersController() (err error, res interface{}) {

	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()), nil
	}

	return nil, users
}

// delete user by id
func (u *UserRepository) DeleteUserController(id int) error {

	result := config.DB.Delete(&models.User{}, id)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "user not found",
		})
	}
	return nil
}

// update user by id
func (u *UserRepository) UpdateUserController(user *models.User, id int) error {

	var users models.User

	result := config.DB.Model(&users).Where("id = ?", id).Updates(&user)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "user not found",
		})
	}
	return nil
}

// login user
func (u *UserRepository) LoginUserController(user *models.User) (error, string) {

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "login failed",
		}), ""
	}

	token, err := middlewares.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "login failed",
			"error":   err.Error(),
		}), ""
	}
	return nil, token
}
