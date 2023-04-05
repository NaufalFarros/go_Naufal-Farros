package Controllers

import (
	// "23_Unit_Testing/Praktikum/RestFullApi/Controllers"
	// "23_Unit_Testing/Praktikum/RestFullApi/Database"
	// "23_Unit_Testing/Praktikum/RestFullApi/Models"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestGetBookController(t *testing.T) {
	// create new echo instance
	e := echo.New()

	// create request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// call controller function
	if assert.NoError(t, Controllers.GetBookController(c)) {
		// check response status code
		assert.Equal(t, http.StatusOK, rec.Code)
		// check response body
		assert.Contains(t, rec.Body.String(), "success get all Books")
	}
}

func TestGetBooksController(t *testing.T) {
	// create new echo instance
	e := echo.New()

	// create new Book instance
	book := Models.Book{
		Title:  "Harry Potter and the Philosopher's Stone",
		Author: "J.K. Rowling",
	}

	// insert new book to database
	Database.DB.Create(&book)

	// create request with book id parameter
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(string(book.ID))

	// call controller function
	if assert.NoError(t, Controllers.GetBooksController(c)) {
		// check response status code
		assert.Equal(t, http.StatusOK, rec.Code)
		// check response body
		assert.Contains(t, rec.Body.String(), "success get Books by id")
		assert.Contains(t, rec.Body.String(), book.Title)
		assert.Contains(t, rec.Body.String(), book.Author)
	}

	// delete book from database
	Database.DB.Delete(&book)
}

func TestCreateBooksController(t *testing.T) {
	// create new echo instance
	e := echo.New()

	// create request body
	requestBody := `{"title":"The Lord of the Rings","author":"J.R.R. Tolkien"}`

	// create request with request body
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// call controller function
	if assert.NoError(t, Controllers.CreateBooksController(c)) {
		// check response status code
		assert.Equal(t, http.StatusOK, rec.Code)
		// check response body
		assert.Contains(t, rec.Body.String(), "success create new Books")
		assert.Contains(t, rec.Body.String(), "The Lord of the Rings")
		assert.Contains(t, rec.Body.String(), "J.R.R. Tolkien")
	}

	// delete created book from database
	var book Models.Book
	Database.DB.Where("title = ?", "The Lord of the Rings").First(&book)
	Database.DB.Delete(&book)
}

func TestDeleteBooksController(t *testing.T) {
	// create new echo instance
	e := echo.New()

	// create new Book instance
	book := Models.Book{
		Title:  "The Hobbit",
		Author: "J.R.R. Tolkien",
	}

	// insert new book to database
	Database.DB.Create(&book)

	// create new request
	req := httptest.NewRequest(http.MethodDelete, "/books/"+strconv.Itoa(int(book.ID)), nil)
	rec := httptest.NewRecorder()

	// set up echo context
	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(book.ID)))

	// call the handler function
	if assert.NoError(t, DeleteBooksController(c)) {
		// check the response status code
		assert.Equal(t, http.StatusOK, rec.Code)

		// check the response body
		expected := `{"message":"success delete Books by id","Books":{"ID":1,"created_at":"2023-04-05T15:22:29.079657+08:00","updated_at":"2023-04-05T15:22:29.079657+08:00","title":"The Hobbit","author":"J.R.R. Tolkien"}}` + "\n"
		assert.Equal(t, expected, rec.Body.String())

		// check if the book is really deleted from the database
		var deletedBook Models.Book
		err := Database.DB.Where("id = ?", book.ID).First(&deletedBook).Error
		if assert.Error(t, err) {
			assert.Equal(t, gorm.ErrRecordNotFound.Error(), err.Error())
		}
	}
}

func TestUpdateBooksController(t *testing.T) {
	// create new echo instance
	e := echo.New()

	// create new Book instance
	book := Models.Book{
		Title:  "The Hobbit",
		Author: "J.R.R. Tolkien",
	}

	// insert new book to database
	Database.DB.Create(&book)

	// update book's title
	book.Title = "The Lord of the Rings"

	// convert book to JSON
	bookJSON, err := json.Marshal(book)
	if err != nil {
		t.Errorf("failed to marshal book to json: %v", err)
	}

	// create new http request
	req := httptest.NewRequest(http.MethodPut, "/books/"+strconv.Itoa(int(book.ID)), bytes.NewBuffer(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// create new echo context
	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(book.ID)))

	// call update books controller
	if err := UpdateBooksController(c); err != nil {
		t.Errorf("failed to update books: %v", err)
	}

	// check http status code
	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	// check response body
	var response struct {
		Message string      `json:"message"`
		Book    Models.Book `json:"Books"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)
	}
	if response.Message != "success update Books by id" {
		t.Errorf("expected message %q but got %q", "success update Books by id", response.Message)
	}
	if response.Book.ID != book.ID {
		t.Errorf("expected book id %d but got %d", book.ID, response.Book.ID)
	}
	if response.Book.Title != "The Lord of the Rings" {
		t.Errorf("expected book title %q but got %q", "The Lord of the Rings", response.Book.Title)
	}
	if response.Book.Author != "J.R.R. Tolkien" {
		t.Errorf("expected book author %q but got %q", "J.R.R. Tolkien", response.Book.Author)
	}
}
