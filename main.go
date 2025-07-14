package main

import (
	"example.com/investment-calulator/db"
	"example.com/investment-calulator/routes"
	"github.com/gin-gonic/gin"
)

func main () {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}

