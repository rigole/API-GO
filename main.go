package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
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

func createBook(c *gin.Context){
	var newBook book

	if err:= c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}


func bookById( c *gin.Context){
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context){
	id, ok := c.GetQuery("id")

	if ok == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Id parameter  is missing."})
		return
	}

	book, err := getBookId(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})	
		return
	}

	if book.Quanity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)

}

func getBookById(id string) (*book, error){
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func main()  {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}