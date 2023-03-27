package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Movie struct {
	Title    string
	Year     string
	Rated    string
	Released string
	Runtime  string
	Genre    string
	Director string
	Writer   string
	Actors   string
	Plot     string
	Language string
	Country  string
	Awards   string
	Poster   string
	Ratings  []struct {
		Source string
		Value  string
	}
	Metascore  string
	ImdbRating string
	ImdbVotes  string
	ImdbID     string
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}

type ListMovie struct {
	Search []struct {
		Title  string
		Year   string
		ImdbID string
		Type   string
		Poster string
	}
	TotalResults string
	Response     string
}

func getIMDBbyID(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)
	url := "https://www.omdbapi.com/?apikey=a818034f&i=" + id

	res, err := http.Get(url)
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}
	defer res.Body.Close()

	var movie Movie
	err = json.NewDecoder(res.Body).Decode(&movie)
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, movie)
}

func getListMoviebyPage(c echo.Context) error {

	search := c.Param("search")
	fmt.Println(search)
	page := c.Param("page")
	fmt.Println(page)

	url := "https://www.omdbapi.com/?apikey=a818034f&s=" + search + "&page=" + page
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}
	defer res.Body.Close()

	fmt.Println(res.Body)

	var movie ListMovie
	err = json.NewDecoder(res.Body).Decode(&movie)
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, movie)

}

func getListMoviebyType(c echo.Context) error {

	typeMovie := c.Param("type")
	search := c.Param("search")
	url := "https://www.omdbapi.com/?apikey=a818034f&type=" + typeMovie + "&s=" + search

	res, err := http.Get(url)

	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}
	defer res.Body.Close()

	var movie ListMovie
	err = json.NewDecoder(res.Body).Decode(&movie)
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, movie)

}

func main() {

	e := echo.New()

	e.GET("/movie/:id", getIMDBbyID)
	e.GET("/movies/:search/:page", getListMoviebyPage)
	e.GET("/movie/:type/:search", getListMoviebyType)
	e.Logger.Fatal(e.Start(":8000"))
}
