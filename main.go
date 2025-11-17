package main

import (
	"OnlineServer/db"
	"OnlineServer/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.InitRoutes(server)

	server.Run(":3000")
}
