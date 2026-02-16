package main

import (
	"example.com/booking-project/db"
	"example.com/booking-project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // initialize the database
	server := gin.Default()

	routes.RegisterRoutes(server) // register the routes

	server.Run(":8080") // localhost
}
