package controllers

import (
	"net/http"

	"example.com/models"
  "example.com/config"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
  DB := config.ConnectDB()
	var products []models.Product

	if err := DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": products})
}
