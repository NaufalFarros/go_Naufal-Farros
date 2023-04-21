package controllers

import (
	"time"

	"github.com/NaufalFarros/golang-fiber/config"
	"github.com/NaufalFarros/golang-fiber/database"
	"github.com/NaufalFarros/golang-fiber/helper"
	"github.com/NaufalFarros/golang-fiber/helper/responseJSON"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         uint      `json:"id" `
	Username   string    `json:"username" validate:"required"`
	Password   string    `json:"password" validate:"required"`
	First_name string    `json:"first_name" validate:"required"`
	Last_name  string    `json:"last_name" validate:"required"`
	RoleId     int       `json:"role_id" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type UpdateUsers struct {
	Password   string `json:"username"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}

// hash password
func HashingPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// compare password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// create user
func CreateUsers(c *fiber.Ctx) error {
	var user = User{}

	if err := c.BodyParser(&user); err != nil {
		return responseJSON.ResponseJSON(c, 400, err.Error(), nil)
	}

	// cek validasi
	errors := helper.ValidationStruct(c, user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// jika username sudah ada
	if database.Database.Db.Where("username = ?", user.Username).First(&user).RowsAffected > 0 {
		return responseJSON.ResponseJSON(c, 400, "username sudah digunakan", nil)
	}

	// hash password lalu simpan ke database
	hash, _ := HashingPassword(user.Password)
	user.Password = hash
	user.Created_at = time.Now()
	user.Updated_at = time.Now()

	database.Database.Db.Create(&user)

	return c.JSON(user)
}

func LoginUsers(c *fiber.Ctx) error {
	var inputuser = User{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}
	var user = User{}

	if err := c.BodyParser(&inputuser); err != nil {
		return responseJSON.ResponseJSON(c, 400, err.Error(), nil)
	}

	// cek username
	if database.Database.Db.Where("username = ?", inputuser.Username).First(&user).RowsAffected == 0 {
		return responseJSON.ResponseJSON(c, 400, "password atau username salah", nil)
	}
	// cek password
	if !CheckPasswordHash(inputuser.Password, user.Password) {
		return responseJSON.ResponseJSON(c, 400, "password atau username salah", nil)
	}

	// expired token
	expiredToken := time.Now().Add(time.Hour * 24)
	claims := &config.JTWClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "NaufalFarros",
			ExpiresAt: jwt.NewNumericDate(expiredToken),
		},
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// sign token
	tokenString, err := token.SignedString(config.JWT_KEY)
	if err != nil {
		return responseJSON.ResponseJSON(c, 400, err.Error(), nil)
	}
	// set token ke cookie

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    tokenString,
		HTTPOnly: true,
		Expires:  expiredToken,
	})

	return responseJSON.ResponseJSON(
		c, 200, "Login berhasil", "",
	)
}

func LogoutUsers(c *fiber.Ctx) error {
	// hapus cookie
	// set token ke cookie

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HTTPOnly: true,
		Expires:  time.Now(),
		MaxAge:   -1,
	})
	return responseJSON.ResponseJSON(c, 200, "Logout berhasil", nil)
}

func GetUsers(c *fiber.Ctx) error {
	// get user yang login cookie
	var user = User{}
	// get token dari cookie
	token := c.Cookies("token")
	// decode token
	claims := &config.JTWClaim{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})
	if err != nil {
		return responseJSON.ResponseJSON(c, 400, err.Error(), nil)
	}
	// get user dari database
	if database.Database.Db.Where("username = ?", claims.Username).First(&user).RowsAffected == 0 {
		return responseJSON.ResponseJSON(c, 400, "user tidak ditemukan", nil)
	}
	return responseJSON.ResponseJSON(c, 200, "user ditemukan", user)

}

func UpdateUser(c *fiber.Ctx) error {
	var user = UpdateUsers{}
	var userdb = User{}

	// get token dari cookie
	token := c.Cookies("token")
	// decode token
	claims := &config.JTWClaim{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})
	if err != nil {
		return responseJSON.ResponseJSON(c, 400, err.Error(), nil)
	}
	// update user
	if err := c.BodyParser(&user); err != nil {
		return responseJSON.ResponseJSON(c, 400, err.Error(), nil)
	}
	// cek parameter yang di update
	if user.Password != " " {
		hash, _ := HashingPassword(user.Password)
		user.Password = hash
	}
	if user.First_name != "  " {
		user.First_name = user.First_name
	}
	if user.Last_name != " " {
		user.Last_name = user.Last_name
	}
	// update user
	database.Database.Db.Model(&userdb).Where("username = ?", claims.Username).Updates(user)

	return responseJSON.ResponseJSON(c, 200, "user berhasil di update", nil)

}
