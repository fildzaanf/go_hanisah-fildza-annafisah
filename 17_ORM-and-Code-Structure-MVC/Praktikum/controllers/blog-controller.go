package controllers

import (
	"Praktikum/config"
	"Praktikum/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all blogs
func GetBlogsController(c echo.Context) error {
	var blogs []models.Blog

	if err := config.DB.Preload("User").Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// get blog by id
func GetBlogController(c echo.Context) error {
	var blog models.Blog

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user ID",
		})
	}

	if err := config.DB.Preload("User").First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "blog not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get blog",
		"blog":     blog,
	})
}

// create new blog
func CreateBlogController(c echo.Context) error {
	blog := models.Blog{}
	var user models.User

	if err := c.Bind(&blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.First(&user, blog.UserId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user ID")
	}

	blog.User = user

	if err := config.DB.Create(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})

}

// delete blog by id
func DeleteBlogController(c echo.Context) error {
	var blog models.Blog

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user ID",
		})
	}

	if err := config.DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "blog not found",
		})
	}

	if err := config.DB.Delete(&models.Blog{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete blog",
	})
}

// update blog by id
func UpdateBlogController(c echo.Context) error {
	var blog models.Blog

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user ID",
		})
	}

	if err := config.DB.Preload("User").First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "blog not found",
		})
	}

	if err := c.Bind(&blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update blog",
		"blog":     blog,
	})
}
