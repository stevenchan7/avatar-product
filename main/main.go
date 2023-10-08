package main

import (
	"net/http"

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
		c.JSON(http.StatusOK, gin.H{"message": "Hello sayang"})
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
		auth.POST("/register", func(c *gin.Context) {
			var userBinder models.RegisterUserInput

			if err := c.ShouldBindJSON(&userBinder); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "err": err.Error()})
				return
			}

			newUser := models.User{Username: userBinder.Username, Password: userBinder.Password}

			if err := DB.Create(&newUser).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "err": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"success": true, "data": newUser})
		})

		auth.POST("/login", func(c *gin.Context) {
			var userBinder models.LoginUserInput

			if err := c.ShouldBind(&userBinder); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "err": err})
			}

			token, err := models.LoginCheck(userBinder.Username, userBinder.Password)

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"success": false, "err": err})
				return
			}

			c.JSON(http.StatusOK, gin.H{"success": true, "token": token})
		})
	}

	r.Run()
}
