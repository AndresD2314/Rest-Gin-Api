package routes

import (
	"github.com/AndresD-2314/Rest-Gin-Api/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/api/getAll", controller.GetAllProducts)
	router.GET("/api/getProduct/:id", controller.GetProductById)
	router.PUT("/api/CreateProduct", controller.CreateProduct)
	router.POST("/api/updateProduct/:id", controller.UpdateProduct)
	router.DELETE("/api/deleteProduct/:id", controller.DeleteProduct)
}
