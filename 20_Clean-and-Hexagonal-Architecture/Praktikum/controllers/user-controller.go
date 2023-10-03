package controllers

import (
	"Praktikum/entity"
	"Praktikum/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUC usecase.UserUseCase
}

func New(userUC usecase.UserUseCase) *UserController {
	return &UserController{
		userUC,
	}
}

func (uc *UserController) CreateUserController(c echo.Context) error {
	user := entity.User{}
	c.Bind(&user)

	err := uc.userUC.CreateUserController(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "create new user success",
		"user":    user,
	})
}

func (uc *UserController) GetUsersController(c echo.Context) error {

	err, res := uc.userUC.GetUsersController()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get all users success",
		"users":   res,
	})
}

func (uc *UserController) LoginUserController(c echo.Context) error {
	user := entity.User{}
	c.Bind(&user)

	err, token := uc.userUC.LoginUserController(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	userResponse := entity.UserResponse{ID: int(user.ID), Name: user.Name, Email: user.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login success",
		"user":    userResponse,
	})
}
