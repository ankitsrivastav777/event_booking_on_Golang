package main

import (
	"booking/rest-api/db"
	"booking/rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // Initialize the database connection
	server := gin.Default()

	routes.RegisterRoutes(server) // Register the routes defined in the routes package
	// Start the server on port 8080
	if err := server.Run(":8080"); err != nil {
		panic(err)
	}
}
