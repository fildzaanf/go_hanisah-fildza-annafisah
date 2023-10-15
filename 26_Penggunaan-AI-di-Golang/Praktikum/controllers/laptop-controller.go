package controllers

import (
	"Praktikum/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func RecommendedLaptopController(c echo.Context) error {

	var input models.Input
	err := c.Bind(&input)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON format")
	}

	content := fmt.Sprintf("Recommended laptop with a budget of %s for %s.", input.Budget, input.Purpose)

	resp, err := RecommendedLaptop(content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status: "Failed",
			Data:   "Failed to Provide Recommendations",
		})
	}

	data := models.Response{
		Status: "Success",
		Data:   resp,
	}

	return c.JSON(http.StatusOK, data)

}
