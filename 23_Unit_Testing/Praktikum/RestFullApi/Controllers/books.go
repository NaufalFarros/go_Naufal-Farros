package Controllers

import (
	"23_Unit_Testing/Praktikum/RestFullApi/Database"
	"23_Unit_Testing/Praktikum/RestFullApi/Models"
	"net/http"

	"github.com/labstack/echo"
)

// get all users
func GetBookController(c echo.Context) error {
	var Books []Models.Book

	if err := Database.DB.Find(&Books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all Books",
		"Books":   Books,
	})
}

// get user by id
func GetBooksController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	var Books Models.Book

	if err := Database.DB.Where("id = ?", id).First(&Books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get Books by id",
		"Books":   Books,
	})
}

// create new user
func CreateBooksController(c echo.Context) error {
	Books := Models.Book{}
	c.Bind(&Books)

	if err := Database.DB.Save(&Books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new Books",
		"Books":   Books,
	})
}

// delete user by id
func DeleteBooksController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	var Books Models.Book

	if err := Database.DB.Where("id = ?", id).Delete(Books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete Books by id",
		"Books":   Books,
	})
}

// update user by id
func UpdateBooksController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	var Books Models.Book

	if err := Database.DB.Where("id = ?", id).First(&Books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&Books)

	if err := Database.DB.Save(&Books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update Books by id",
		"Books":   Books,
	})
}
