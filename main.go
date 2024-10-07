package main

import (
	"example/mysql-api/database"
	"example/mysql-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDb()
	defer database.CloseDb()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run()
}
