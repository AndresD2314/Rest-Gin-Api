package main

import (
	"github.com/AndresD-2314/Rest-Gin-Api/database"
	"github.com/AndresD-2314/Rest-Gin-Api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	database.ConnectDB()

	routes.UserRoute(router)

	router.Run(":8080")

}
