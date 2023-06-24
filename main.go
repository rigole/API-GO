package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
)

type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

var books = []book{
	{ID: "1",Title: "Getting things done", Author: "David Allen", Quantity: 5},
	{ID: "2",Title: "The millionaire Fastlane", Author: "MJ DeMarco", Quantity: 14},
	{ID: "3",Title: "Clean code ", Author: "Martin Cecil", Quantity: 10},
}

func getBooks(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, books)
}

func main()  {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}