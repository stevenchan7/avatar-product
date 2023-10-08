package main

import (
	"net/http"
	"os"

	"example.com/config"
	"example.com/controllers"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Call ConnectDB
	DB := config.ConnectDB()

	r.GET("/", func(c *gin.Context) {
		secret := os.Getenv("TOKEN_SECRET")
		c.JSON(http.StatusOK, gin.H{"message": "Hello sayang", "Secret": secret})
	})

	r.GET("/products", func(c *gin.Context) {
		var products []models.Product

		if err := DB.Find(&products).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "data": products})
	})

	admin := r.Group("/admin")
	{
		admin.POST("/add-product", controllers.PostProduct)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)

		auth.POST("/login", controllers.Login)
	}

	r.Run()
}
