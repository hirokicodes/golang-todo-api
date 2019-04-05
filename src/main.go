package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func main() {
	fmt.Println("hello world")

	r := gin.Default()
	r.GET("/", homePage)

	r.Run()
}
