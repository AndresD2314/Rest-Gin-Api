package controller

import (
	"github.com/AndresD-2314/Rest-Gin-Api/database"
	"github.com/AndresD-2314/Rest-Gin-Api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllProducts(c *gin.Context) {
	var products []models.Producto
	database.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	var product models.Producto

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
	}

	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product models.Producto

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el producto"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Producto

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de producto inválidos"})
		return
	}

	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el producto"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Producto

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Producto no ha sido encontrado"})
		return
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "No se pudo eliminar el producto desado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Producto eliminado de manera correcta"})
}
