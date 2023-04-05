package Controllers

import (
	"23_Unit_Testing/Praktikum/RestFullApi/Models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestLoginController(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody, _ := json.Marshal(map[string]string{
		"email":    "user@example.com",
		"password": "password",
	})
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, LoginController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "token")
	}
}

func TestGetUsersController(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserController(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, GetUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateUserController(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody, _ := json.Marshal(map[string]string{
		"name":     "John Doe",
		"email":    "johndoe@example.com",
		"password": "password",
	})
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, CreateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteUserController(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateUserController(t *testing.T) {
	// Setup
	Database.Connect()
	e := Controllers.NewEcho()
	defer Database.DB.Close()

	// Create dummy user
	user := Models.User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "password",
	}
	Database.DB.Create(&user)

	// Update user data
	reqBody, _ := json.Marshal(map[string]string{
		"name":     "Jane Doe",
		"email":    "janedoe@example.com",
		"password": "newpassword",
	})
	req := httptest.NewRequest(http.MethodPut, "/users/"+user.ID.String(), bytes.NewBuffer(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues(user.ID.String())

	// Assertions
	if assert.NoError(t, Controllers.UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"message":"success update user by id"`)
		assert.Contains(t, rec.Body.String(), `"name":"Jane Doe"`)
		assert.Contains(t, rec.Body.String(), `"email":"janedoe@example.com"`)
		assert.NotContains(t, rec.Body.String(), `"password":"newpassword"`)
	}
}
