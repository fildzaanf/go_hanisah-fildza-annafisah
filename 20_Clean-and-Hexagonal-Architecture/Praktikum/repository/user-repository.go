package repository

import (
	"Praktikum/config"
	"Praktikum/entity"
	"Praktikum/middlewares"
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	CreateUserController(user *entity.User) error
	GetUsersController() (err error, res interface{})
	LoginUserController(user *entity.User) (error, string)
}

type UserRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepo {

	return &UserRepo{
		db,
	}
}

func (u *UserRepo) CreateUserController(user *entity.User) error {

	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("All fields must be filled")
	}

	if err := u.db.Save(&user); err != nil {
		return err.Error
	}

	return nil
}

func (u *UserRepo) GetUsersController() (err error, res interface{}) {

	var users []entity.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()), nil
	}

	return nil, users
}

func (u *UserRepo) LoginUserController(user *entity.User) (error, string) {

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "login failed",
		}), ""

	}

	token, err := middlewares.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "login  failed",
			"error":   err.Error(),
		}), ""
	}

	return nil, token
}
