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

	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for i, blog := range blogs {

		var user models.User

		if err := config.DB.First(&user, blog.UserId).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user data")
		}
		blogs[i].User = models.User{Name: user.Name}
	}

	var responsBlogs []map[string]interface{}
	for _, blog := range blogs {

		user := map[string]interface{}{
			"name": blog.User.Name,
		}

		responsBlogs = append(responsBlogs, map[string]interface{}{
			"ID":      blog.ID,
			"Title":   blog.Title,
			"Content": blog.Content,
			"User":    user,
		})
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

	if err := config.DB.First(&blog, id).Error; err != nil {
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

	if err := c.Bind(blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.First(&user, blog.UserId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user ID")
	}

	if err := config.DB.Create(blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	username := user.Name
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog": map[string]interface{}{
			"title":   blog.Title,
			"content": blog.Content,
			"user": map[string]interface{}{
				"name": username,
			},
		},
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
	var blogs models.Blog
	blog := new(models.Blog)

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user ID",
		})
	}

	if err := c.Bind(blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.First(&blogs, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "blog not found",
		})
	}

	blogs.Title = blog.Title
	blogs.Content = blog.Content

	if err := config.DB.Save(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update blog",
		"blog":     blogs,
	})
}
