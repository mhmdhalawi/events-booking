package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mhmdhalawi/events-booking/db"
	"github.com/mhmdhalawi/events-booking/routes"
)

func main() {

	db.InitDB()

	server := gin.Default()
	server.SetTrustedProxies([]string{"localhost"})

	router := server.Group("/")
	routes.AddRoutes(router)

	server.Run("localhost:9000")

}
