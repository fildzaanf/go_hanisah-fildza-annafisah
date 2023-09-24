package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {

	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user ID",
		})
	}

	for i := range users {
		if users[i].Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success get user by ID",
				"user":    &users[i],
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {

	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user ID",
		})
	}

	for i := range users {
		if users[i].Id == id {

			users = append(users[:i], users[i+1:]...)

			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "user deleted successfully",
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {

	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user ID",
		})
	}

	for i := range users {
		if users[i].Id == id {
			updatedUser := new(User)
			if err := c.Bind(updatedUser); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": "invalid request data",
				})
			}

			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			users[i].Password = updatedUser.Password

			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "user updated successfully",
				"user":    &users[i],
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()

	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users/add", CreateUserController)
	e.PUT("/users/update/:id", UpdateUserController)
	e.DELETE("/users/delete/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
