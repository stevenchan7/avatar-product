package main

import (
	"net/http"

	"example.com/hello"
	// "example.com/salam/models"
	"github.com/gin-gonic/gin"
)

func main() {
	hello.Hello()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello sayang"})
	})

	r.Run()
}
