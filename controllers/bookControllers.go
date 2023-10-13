package controllers

import (
	"net/http"
	"strconv"
	"tugas/configs"
	"tugas/models"

	"github.com/labstack/echo/v4"
)

// get all users
func GetBooksController(c echo.Context) error {
	var books []models.Book

	if err := configs.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"books":   books,
	})

}

// get user by id
func GetBooksByIDController(c echo.Context) error {
	id := c.Param("id")
	var books []models.Book
	err := configs.DB.First(&books, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get id users",
		"books":   books,
	})
}

// create new user
func CreateBooksController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := configs.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"book":    book,
	})
}

// delete user by id
func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var books []models.Book
	eror := configs.DB.Delete(&books, id)
	if eror != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete id books",
		"books":   books,
	})
}

// update user by id
func UpdateBookController(c echo.Context) error {
	var books, updateData models.Book
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&updateData)
	result := configs.DB.First(&books, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error",
		})
	}
	configs.DB.Model(&books).Updates(updateData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete id users",
		"books":   books,
	})
}
