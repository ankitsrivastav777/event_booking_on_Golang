package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)       // Create a new event
	server.GET("/events/:id", getEvent)       // Get a specific event by ID
	server.PUT("/events/:id", updateEvent)    // Update an existing event by ID
	server.DELETE("/events/:id", deleteEvent) // Delete an event by ID
	server.POST("/signup", signup)            // User signup route
}
