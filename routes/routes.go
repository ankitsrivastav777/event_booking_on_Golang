package routes

import (
	"booking/rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)                      // Apply the authentication middleware to all routes in this group
	authenticated.POST("/events", createEvent)                       // Create a new event
	authenticated.PUT("/events/:id", updateEvent)                    // Update an existing event by ID
	authenticated.DELETE("/events/:id", deleteEvent)                 // Delete an event by ID
	authenticated.POST("/events/:id/register", registerForEvent)     // Book an event by ID
	authenticated.DELETE("/events/:id/register", CancelRegistration) // Unregister from an event by ID
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // Get a specific event by ID
	server.POST("/signup", signup)      // User signup route
	server.POST("/login", login)        // User login route

}
