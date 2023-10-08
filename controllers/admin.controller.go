package controllers

import(
  "net/http"

  "github.com/gin-gonic/gin"
  "example.com/models"
  "example.com/config"
)

func PostProduct(c *gin.Context) {
  // Call ConnectDB
	DB := config.ConnectDB()

  var productInput models.ProductInput

  // Bind input and check error at the same time
  if err := c.ShouldBindJSON(&productInput); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
  }

  newProduct := models.Product{Title: productInput.Title, Desc: productInput.Desc, Image: productInput.Image, Playstore: productInput.Playstore, Appstote: productInput.Appstote}

  // Insert new product to database
  if err := DB.Create(&newProduct).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"success": true, "data": newProduct})
}