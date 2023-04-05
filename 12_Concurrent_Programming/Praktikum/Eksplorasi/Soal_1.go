package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type dataResponse []struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      struct {
		Rate  float64 `json:"rate"`
		Count int     `json:"count"`
	} `json:"rating"`
}

func fetchAPI(url string, channel chan dataResponse) {
	resp, err := http.Get(url)
	if err != nil {
		channel <- dataResponse{
			{
				ID:          0,
				Title:       "Error",
				Price:       0,
				Description: "Error",
				Category:    "Error",
				Image:       "Error",
				Rating: struct {
					Rate  float64 "json:\"rate\""
					Count int     "json:\"count\""
				}{Rate: 0, Count: 0},
			},
		}
		return
	}
	defer resp.Body.Close()

	var data dataResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}

	channel <- data
}

func main() {

	ch := make(chan dataResponse)

	go fetchAPI("https://fakestoreapi.com/products", ch)

	fmt.Println("Product List:")
	for _, product := range <-ch {
		fmt.Println("=============")
		fmt.Println("Title :", product.Title)
		fmt.Println("Price :", product.Price)
		fmt.Println("Description :", product.Description)
		fmt.Println("Image :", product.Image)
		fmt.Println("Category :", product.Category)
		fmt.Println("Rating :", product.Rating.Rate)
		fmt.Println("Rating Count :", product.Rating.Count)
		fmt.Println("=============")
	}

	// fmt.Println(<-ch)

}
