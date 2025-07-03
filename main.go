package main

import (
	"api-login/config"
	"api-login/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080")
}
