package controllers

import (
	"net/http"

	"example.com/config"
	"example.com/models"
	"github.com/gin-gonic/gin"
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

func EditProduct(c *gin.Context) {
	DB := config.ConnectDB()
	// Bind user input
	var productInput models.EditProductInput

	if err := c.ShouldBind(&productInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "err": err.Error()})
	}

	// Find product in database
	prodID := productInput.ID

	var updatedProduct models.Product

	DB.First(&updatedProduct, "id = ?", prodID)

	// Update product in database
	// updatedProduct.UpdatedAt = time.Now()
	DB.Model(&updatedProduct).Updates(productInput)

	c.JSON(http.StatusOK, gin.H{"success": true, "data": updatedProduct})
}
