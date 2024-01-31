package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Think and grow rich", Author: "napalen hill", Quantity: 3},
	{ID: "3", Title: "Atomic Habits", Author: "james clear", Quantity: 5},
	{ID: "2", Title: "nothing", Author: "dale carniegie", Quantity: 4},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)

	router.Run("localhost:3000")
}
