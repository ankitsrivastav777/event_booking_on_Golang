package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Define a simple route
	server.GET("/events", getEvents)

	// Start the server on port 8080
	if err := server.Run(":8080"); err != nil {
		panic(err)
	}
}

func getEvents(context *gin.Context) {
	// This function will handle the /events route
	context.JSON(http.StatusOK, gin.H{"message": "Hi Events!"})

}
