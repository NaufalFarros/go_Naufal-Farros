package Controllers

import (
	"net/http"

	"22_Middleware/Praktikum/Models"

	"22_Middleware/Praktikum/Database"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type loginResponse struct {
	Token string `json:"token"`
}

func LoginController(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Validate user credentials here
	var user Models.User
	if err := Database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Compare password

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	t, err := token.SignedString([]byte("NaufalFarros"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, loginResponse{Token: t})
}

// get all users
func GetUsersController(c echo.Context) error {
	var users []Models.User

	if err := Database.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	var user Models.User

	if err := Database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := Models.User{}
	c.Bind(&user)

	// hash password
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user.Password = string(hashedPassword)

	if err := Database.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	var user Models.User

	if err := Database.DB.Where("id = ?", id).Delete(user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by id",
		"user":    user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	var user Models.User

	if err := Database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&user)

	// hash password
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user.Password = string(hashedPassword)

	if err := Database.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user by id",
		"user":    user,
	})
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
