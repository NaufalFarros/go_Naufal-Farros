package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {

	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Book struct {
	gorm.Model
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Book{})
}

// get all users
func GetBookController(c echo.Context) error {
	var Books []Book

	if err := DB.Find(&Books).Error; err != nil {
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
	var Books Book

	if err := DB.Where("id = ?", id).First(&Books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get Books by id",
		"Books":   Books,
	})
}

// create new user
func CreateBooksController(c echo.Context) error {
	Books := Book{}
	c.Bind(&Books)

	if err := DB.Save(&Books).Error; err != nil {
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
	var Books Book

	if err := DB.Where("id = ?", id).Delete(Books).Error; err != nil {
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
	var Books Book

	if err := DB.Where("id = ?", id).First(&Books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&Books)

	if err := DB.Save(&Books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update Books by id",
		"Books":   Books,
	})
}

func main() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/books", GetBookController)
	e.GET("/books/:id", GetBooksController)
	e.POST("/books", CreateBooksController)
	e.DELETE("/books/:id", DeleteBooksController)
	e.PUT("/books/:id", UpdateBooksController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
