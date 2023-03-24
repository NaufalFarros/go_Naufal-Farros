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

var users = []User{
	{
		Id:       1,
		Name:     "John Doe",
		Email:    "johnDoe@gmail.com",
		Password: "123456",
	},
	{
		Id:       2,
		Name:     "Jane Doe",
		Email:    "Janedoe@gmail.com",
		Password: "123456",
	},
	{
		Id:       3,
		Name:     "John Smith",
		Email:    "JohnSmit@gmail.com",
		Password: "123456",
	},
	{
		Id:       4,
		Name:     "Jane Smith",
		Email:    "JaneSMith@gmail.com",
		Password: "123456",
	},
	{
		Id:       5,
		Name:     "Naufal Farros",
		Email:    "naufalfarros05@gmail.com",
		Password: "123456",
	},
}

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUsersControllers(c echo.Context) error {
	user, _ := strconv.Atoi(c.Param("id"))
	for _, value := range users {
		if value.Id == user {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get user",
				"user":     value,
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	user, _ := strconv.Atoi(c.Param("id"))

	for index, value := range users {
		if value.Id == user {
			users = append(users[:index], users[index+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success delete user",
				"user":     users,
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found, failed to delete",
	})

}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	user, _ := strconv.Atoi(c.Param("id"))

	for index, value := range users {
		if value.Id == user {
			// cek apakah parameter yang dikirimkan ada atau tidak
			if c.FormValue("name") == "" {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"messages": "name is required",
				})
			}
			if c.FormValue("email") == "" {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"messages": "email is required",
				})
			}
			if c.FormValue("password") == "" {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"messages": "password is required",
				})
			}

			users[index].Name = c.FormValue("name")
			users[index].Email = c.FormValue("email")
			users[index].Password = c.FormValue("password")

			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success update user",
				"user":     users[index],
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "user not found, failed to update",
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
	e.GET("/users/:id", GetUsersControllers)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.POST("/users", CreateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
