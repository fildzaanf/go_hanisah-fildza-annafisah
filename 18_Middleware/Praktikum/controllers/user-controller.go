package controllers

import (
	"Praktikum/config"
	"Praktikum/models"
	"Praktikum/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	var user models.User

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user ID",
		})
	}

	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "user not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user",
		"user":     user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := &models.User{}

	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	var users models.User

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user ID",
		})
	}

	if err := config.DB.First(&users, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "user not found",
		})
	}

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	var users models.User
	user := new(models.User)

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user ID",
		})
	}

	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.First(&users, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	users.Name = user.Name
	users.Email = user.Email
	users.Password = user.Password

	if err := config.DB.Save(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updated user",
		"user":    users,
	})
}

// login user
func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "login failed",
			"error":   err.Error(),
		})
	}

	token, err := middlewares.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "login failed",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserResponse{ID: int(user.ID), Name: user.Name, Email: user.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user login successfully",
		"user":    userResponse,
	})
}